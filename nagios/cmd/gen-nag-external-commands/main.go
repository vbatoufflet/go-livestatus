package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

var out = flag.String("o", "", "output file to write to")

var nagCmdIndex = "https://www.nagios.org/developerinfo/externalcommands/"
var nagCmdURLRgxp = regexp.MustCompile(`<a href='(commandinfo.php\?command_id=[0-9]+)'>([A-Z_]+)</a></td></tr>`)

type nagCmd struct {
	def  string
	desc string

	name string
	args []string
}

var nagCmds = []nagCmd{}

type argDef struct {
	t      string
	fmtStr string
}

var argTypes = map[string]argDef{
	"host_name":               {"string", `Hostname(c, %s)`},
	"service_description":     {"string", `ServiceDescription(c, %s)`},
	"sticky":                  {"bool", `Sticky(c, %s)`},
	"notify":                  {"bool", `Notify(c, %s)`},
	"fixed":                   {"bool", `Fixed(c, %s)`},
	"persistent":              {"bool", "Persistent(c, %s)"},
	"author":                  {"string", `Author(c, %s)`},
	"contact_name":            {"string", `ContactName(c, %s)`},
	"contactgroup_name":       {"string", `ContactGroupName(c, %s)`},
	"hostgroup_name":          {"string", `HostGroupName(c, %s)`},
	"servicegroup_name":       {"string", `ServiceGroupName(c, %s)`},
	"comment":                 {"string", `Comment(c, %s)`},
	"start_time":              {"time.Time", `Start(c, %s)`},
	"end_time":                {"time.Time", `End(c, %s)`},
	"check_time":              {"time.Time", `CheckTime(c, %s)`},
	"notification_time":       {"time.Time", `NotificationTime(c, %s)`},
	"notification_timeperiod": {"string", `NotificationTimePeriod(c, %s)`},
	"duration":                {"time.Duration", `Duration(c, %s)`},
	"trigger_id":              {"int", `TriggerID(c, %s)`},
	"downtime_id":             {"int", `DowntimeID(c, %s)`},
	"comment_id":              {"int", `CommentID(c, %s)`},
	"options":                 {"int", `Options(c, %s)`},
	"value":                   {"string", `Value(c, %s)`},
	"varname":                 {"string", `VarName(c, %s)`},
	"varvalue":                {"string", `VarValue(c, %s)`},
	"event_handler_command":   {"string", `EventHandlerCommand(c, %s)`},
	"check_command":           {"string", `CheckCommand(c, %s)`},
	"timeperiod":              {"string", `TimePeriod(c, %s)`},
	"check_timeperod":         {"string", `CheckTimePeriod(c, %s)`},
	"check_timeperiod":        {"string", `CheckTimePeriod(c, %s)`},
	"check_attempts":          {"int", `CheckAttempts(c, %s)`},
	"check_interval":          {"time.Duration", `Duration(c, %s)`},
	"file_name":               {"string", `FileName(c, %s)`},
	"delete":                  {"bool", `Delete(c, %s)`},
	"status_code":             {"int", `StatusCode(c, %s)`},
	"return_code":             {"int", `ReturnCode(c, %s)`},
	"plugin_output":           {"string", `PluginOutput(c, %s)`},
	"notification_number":     {"int", `NotificationNumber(c, %s)`},
}

func main() {
	flag.Parse()

	var err error
	var file io.WriteCloser

	file = os.Stdout
	if *out != "" {
		file, err = os.Create(*out)
	}
	if err != nil {
		log.Fatalf("error creating file, %v", err.Error())
	}

	defer file.Close()

	client := http.Client{}
	buf := &bytes.Buffer{}
	resp, err := client.Get(nagCmdIndex)
	if err != nil {
		log.Fatalf("error fetching index, %v", err.Error())
	}

	io.Copy(buf, resp.Body)

	cs := nagCmdURLRgxp.FindAllStringSubmatch(buf.String(), -1)
	for _, cUrl := range cs {
		cmd := cUrl[2]
		ref, err := url.Parse(cUrl[1])
		if err != nil {
			log.Println("URL reference not valid, ", err.Error())
			continue
		}

		url := resp.Request.URL.ResolveReference(ref)
		cresp, err := client.Get(url.String())
		if err != nil {
			log.Fatalf("error fetching command %v, %v", cmd, err.Error())
		}

		doc, err := html.Parse(cresp.Body)
		if err != nil {
			log.Println("error parsing html, ", err.Error())
		}
		def, desc := findDefAndDesc(doc)
		nagCmds = append(nagCmds, nagCmd{def, desc, "", nil})
	}

	genCode(file, "nagios", nagCmds)
}

func findDefAndDesc(n *html.Node) (string, string) {
	def := ""
	desc := ""
	foundDefHeader := false
	foundDescHeader := false

	var walk func(n *html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.TextNode && n.Data == "Command Format:" {
			foundDefHeader = true
			return
		}
		if n.Type == html.TextNode && foundDefHeader && def == "" && strings.TrimSpace(n.Data) != "" {
			def = strings.TrimSpace(n.Data)
			return
		}

		if n.Type == html.TextNode && n.Data == "Description:" {
			foundDescHeader = true
			return
		}
		if n.Type == html.TextNode && foundDescHeader && desc == "" && strings.TrimSpace(n.Data) != "" {
			desc = strings.TrimSpace(n.Data)
			desc = strings.Replace(desc, "\n", " ", -1)
			return
		}

		if def != "" && desc != "" {
			return
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return def, desc
}

func genCode(w io.Writer, pkg string, cmds []nagCmd) {
	fmt.Fprintf(w, "package %v\n\n", pkg)
	fmt.Fprintf(w, "import \"time\"\n\n")
	fmt.Fprintf(w, "import lvst \"github.com/vbatoufflet/go-livestatus\"\n\n")

	done := map[string]bool{}
	for _, c := range cmds {
		err := c.parseCommandDef()
		if err != nil {
			continue
		}
		if _, ok := done[c.name]; ok {
			//repeated command, CHANGE_HOST_CHECK_TIMEPERIOD is in there twice.
			continue
		}
		nps := strings.Split(c.name, "_")
		goname := ""
		for _, np := range nps {
			goname = goname + strings.ToUpper(np[0:1]) + strings.ToLower(np[1:])
		}
		fmt.Fprintf(w, "// %v is generated from the nagios external command definition:\n", goname)
		fmt.Fprintf(w, "// Desfinition: %v\n", c.def)
		fmt.Fprintln(w, "// Description:")
		fmt.Fprintf(w, "//  %v\n", c.desc)
		args := []string{}
		for _, a := range c.args {
			args = append(args, fmt.Sprintf("%s %s", a, argTypes[a].t))
		}
		fmt.Fprintf(w, "func %v(%s) lvst.CommandOpFunc{\n", goname, strings.Join(args, ", "))
		fmt.Fprint(w, "\treturn func (c *lvst.Command) {\n")
		fmt.Fprintf(w, "\t\tc.Raw(\"%v\")\n", c.name)
		for _, a := range c.args {
			fmt.Fprintf(w, "\t\t%s\n", fmt.Sprintf(argTypes[a].fmtStr, a))
		}
		fmt.Fprint(w, "\t}\n")
		fmt.Fprint(w, "}\n\n")
		done[c.name] = true
	}
}

var cmdTmpl = regexp.MustCompile("^[A-Z_]+$")
var argTmpl = regexp.MustCompile("^<([a-z_]+)>$")

func (n *nagCmd) parseCommandDef() error {
	parts := strings.Split(n.def, ";")
	if !cmdTmpl.MatchString(parts[0]) {
		return fmt.Errorf("malformed command definition")
	}
	cmd := parts[0]

	args := []string{}
	for _, a := range parts[1:] {
		ms := argTmpl.FindStringSubmatch(a)
		if ms == nil {
			return fmt.Errorf("malformed command definition")
		}
		args = append(args, ms[1])
	}

	n.name = cmd
	n.args = args

	return nil
}
