package nagios

import "time"

import lvst "github.com/vbatoufflet/go-livestatus"

// AcknowledgeHostProblem is generated from the nagios external command definition:
// Desfinition: ACKNOWLEDGE_HOST_PROBLEM;<host_name>;<sticky>;<notify>;<persistent>;<author>;<comment>
// Description:
//  Allows you to acknowledge the current problem for the specified host.  By acknowledging the current problem, future notifications (for the same host state) are disabled.  If the "sticky" option is set to two (2), the acknowledgement will remain until the host returns to an UP state.  Otherwise the acknowledgement will automatically be removed when the host changes state.  If the "notify" option is set to one (1), a notification will be sent out to contacts indicating that the current host problem has been acknowledged.  If the "persistent" option is set to one (1), the comment associated with the acknowledgement will survive across restarts of the Nagios process.  If not, the comment will be deleted the next time Nagios restarts.
func AcknowledgeHostProblem(host_name string, sticky bool, notify bool, persistent bool, author string, comment string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ACKNOWLEDGE_HOST_PROBLEM")
		Hostname(c, host_name)
		Sticky(c, sticky)
		Notify(c, notify)
		Persistent(c, persistent)
		Author(c, author)
		Comment(c, comment)
	}
}

// AcknowledgeSvcProblem is generated from the nagios external command definition:
// Desfinition: ACKNOWLEDGE_SVC_PROBLEM;<host_name>;<service_description>;<sticky>;<notify>;<persistent>;<author>;<comment>
// Description:
//  Allows you to acknowledge the current problem for the specified service.  By acknowledging the current problem, future notifications (for the same servicestate) are disabled.  If the "sticky" option is set to two (2), the acknowledgement will remain until the service returns to an OK state.  Otherwise the acknowledgement will automatically be removed when the service changes state.  If the "notify" option is set to one (1), a notification will be sent out to contacts indicating that the current service problem has been acknowledged.  If the "persistent" option is set to one (1), the comment associated with the acknowledgement will survive across restarts of the Nagios process.  If not, the comment will be deleted the next time Nagios restarts.
func AcknowledgeSvcProblem(host_name string, service_description string, sticky bool, notify bool, persistent bool, author string, comment string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ACKNOWLEDGE_SVC_PROBLEM")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		Sticky(c, sticky)
		Notify(c, notify)
		Persistent(c, persistent)
		Author(c, author)
		Comment(c, comment)
	}
}

// AddHostComment is generated from the nagios external command definition:
// Desfinition: ADD_HOST_COMMENT;<host_name>;<persistent>;<author>;<comment>
// Description:
//  Adds a comment to a particular host.  If the "persistent" field is set to zero (0), the comment will be deleted the next time Nagios is restarted.  Otherwise, the comment will persist across program restarts until it is deleted manually.
func AddHostComment(host_name string, persistent bool, author string, comment string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ADD_HOST_COMMENT")
		Hostname(c, host_name)
		Persistent(c, persistent)
		Author(c, author)
		Comment(c, comment)
	}
}

// AddSvcComment is generated from the nagios external command definition:
// Desfinition: ADD_SVC_COMMENT;<host_name>;<service_description>;<persistent>;<author>;<comment>
// Description:
//  Adds a comment to a particular service.  If the "persistent" field is set to zero (0), the comment will be deleted the next time Nagios is restarted.  Otherwise, the comment will persist across program restarts until it is deleted manually.
func AddSvcComment(host_name string, service_description string, persistent bool, author string, comment string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ADD_SVC_COMMENT")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		Persistent(c, persistent)
		Author(c, author)
		Comment(c, comment)
	}
}

// ChangeContactHostNotificationTimeperiod is generated from the nagios external command definition:
// Desfinition: CHANGE_CONTACT_HOST_NOTIFICATION_TIMEPERIOD;<contact_name>;<notification_timeperiod>
// Description:
//  Changes the host notification timeperiod for a particular contact to what is specified by the "notification_timeperiod" option.  The "notification_timeperiod" option should be the short name of the timeperiod that is to be used as the contact's host notification timeperiod.  The timeperiod must have been configured in Nagios before it was last (re)started.
func ChangeContactHostNotificationTimeperiod(contact_name string, notification_timeperiod string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_CONTACT_HOST_NOTIFICATION_TIMEPERIOD")
		ContactName(c, contact_name)
		NotificationTimePeriod(c, notification_timeperiod)
	}
}

// ChangeContactModattr is generated from the nagios external command definition:
// Desfinition: CHANGE_CONTACT_MODATTR;<contact_name>;<value>
// Description:
//  This command changes the modified attributes value for the specified contact.  Modified attributes values are used by Nagios to determine which object properties should be retained across program restarts.  Thus, modifying the value of the attributes can affect data retention.  This is an advanced option and should only be used by people who are intimately familiar with the data retention logic in Nagios.
func ChangeContactModattr(contact_name string, value string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_CONTACT_MODATTR")
		ContactName(c, contact_name)
		Value(c, value)
	}
}

// ChangeContactModhattr is generated from the nagios external command definition:
// Desfinition: CHANGE_CONTACT_MODHATTR;<contact_name>;<value>
// Description:
//  This command changes the modified host attributes value for the specified contact.  Modified attributes values are used by Nagios to determine which object properties should be retained across program restarts.  Thus, modifying the value of the attributes can affect data retention.  This is an advanced option and should only be used by people who are intimately familiar with the data retention logic in Nagios.
func ChangeContactModhattr(contact_name string, value string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_CONTACT_MODHATTR")
		ContactName(c, contact_name)
		Value(c, value)
	}
}

// ChangeContactModsattr is generated from the nagios external command definition:
// Desfinition: CHANGE_CONTACT_MODSATTR;<contact_name>;<value>
// Description:
//  This command changes the modified service attributes value for the specified contact.  Modified attributes values are used by Nagios to determine which object properties should be retained across program restarts.  Thus, modifying the value of the attributes can affect data retention.  This is an advanced option and should only be used by people who are intimately familiar with the data retention logic in Nagios.
func ChangeContactModsattr(contact_name string, value string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_CONTACT_MODSATTR")
		ContactName(c, contact_name)
		Value(c, value)
	}
}

// ChangeContactSvcNotificationTimeperiod is generated from the nagios external command definition:
// Desfinition: CHANGE_CONTACT_SVC_NOTIFICATION_TIMEPERIOD;<contact_name>;<notification_timeperiod>
// Description:
//  Changes the service notification timeperiod for a particular contact to what is specified by the "notification_timeperiod" option.  The "notification_timeperiod" option should be the short name of the timeperiod that is to be used as the contact's service notification timeperiod.  The timeperiod must have been configured in Nagios before it was last (re)started.
func ChangeContactSvcNotificationTimeperiod(contact_name string, notification_timeperiod string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_CONTACT_SVC_NOTIFICATION_TIMEPERIOD")
		ContactName(c, contact_name)
		NotificationTimePeriod(c, notification_timeperiod)
	}
}

// ChangeCustomContactVar is generated from the nagios external command definition:
// Desfinition: CHANGE_CUSTOM_CONTACT_VAR;<contact_name>;<varname>;<varvalue>
// Description:
//  Changes the value of a custom contact variable.
func ChangeCustomContactVar(contact_name string, varname string, varvalue string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_CUSTOM_CONTACT_VAR")
		ContactName(c, contact_name)
		VarName(c, varname)
		VarValue(c, varvalue)
	}
}

// ChangeCustomHostVar is generated from the nagios external command definition:
// Desfinition: CHANGE_CUSTOM_HOST_VAR;<host_name>;<varname>;<varvalue>
// Description:
//  Changes the value of a custom host variable.
func ChangeCustomHostVar(host_name string, varname string, varvalue string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_CUSTOM_HOST_VAR")
		Hostname(c, host_name)
		VarName(c, varname)
		VarValue(c, varvalue)
	}
}

// ChangeCustomSvcVar is generated from the nagios external command definition:
// Desfinition: CHANGE_CUSTOM_SVC_VAR;<host_name>;<service_description>;<varname>;<varvalue>
// Description:
//  Changes the value of a custom service variable.
func ChangeCustomSvcVar(host_name string, service_description string, varname string, varvalue string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_CUSTOM_SVC_VAR")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		VarName(c, varname)
		VarValue(c, varvalue)
	}
}

// ChangeGlobalHostEventHandler is generated from the nagios external command definition:
// Desfinition: CHANGE_GLOBAL_HOST_EVENT_HANDLER;<event_handler_command>
// Description:
//  Changes the global host event handler command to be that specified by the "event_handler_command" option.  The "event_handler_command" option specifies the short name of the command that should be used as the new host event handler.  The command must have been configured in Nagios before it was last (re)started.
func ChangeGlobalHostEventHandler(event_handler_command string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_GLOBAL_HOST_EVENT_HANDLER")
		EventHandlerCommand(c, event_handler_command)
	}
}

// ChangeGlobalSvcEventHandler is generated from the nagios external command definition:
// Desfinition: CHANGE_GLOBAL_SVC_EVENT_HANDLER;<event_handler_command>
// Description:
//  Changes the global service event handler command to be that specified by the "event_handler_command" option.  The "event_handler_command" option specifies the short name of the command that should be used as the new service event handler.  The command must have been configured in Nagios before it was last (re)started.
func ChangeGlobalSvcEventHandler(event_handler_command string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_GLOBAL_SVC_EVENT_HANDLER")
		EventHandlerCommand(c, event_handler_command)
	}
}

// ChangeHostCheckCommand is generated from the nagios external command definition:
// Desfinition: CHANGE_HOST_CHECK_COMMAND;<host_name>;<check_command>
// Description:
//  Changes the check command for a particular host to be that specified by the "check_command" option.  The "check_command" option specifies the short name of the command that should be used as the new host check command.  The command must have been configured in Nagios before it was last (re)started.
func ChangeHostCheckCommand(host_name string, check_command string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_HOST_CHECK_COMMAND")
		Hostname(c, host_name)
		CheckCommand(c, check_command)
	}
}

// ChangeHostCheckTimeperiod is generated from the nagios external command definition:
// Desfinition: CHANGE_HOST_CHECK_TIMEPERIOD;<host_name>;<timeperiod>
// Description:
//  Changes the valid check period for the specified host.
func ChangeHostCheckTimeperiod(host_name string, timeperiod string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_HOST_CHECK_TIMEPERIOD")
		Hostname(c, host_name)
		TimePeriod(c, timeperiod)
	}
}

// ChangeHostEventHandler is generated from the nagios external command definition:
// Desfinition: CHANGE_HOST_EVENT_HANDLER;<host_name>;<event_handler_command>
// Description:
//  Changes the event handler command for a particular host to be that specified by the "event_handler_command" option.  The "event_handler_command" option specifies the short name of the command that should be used as the new host event handler.  The command must have been configured in Nagios before it was last (re)started.
func ChangeHostEventHandler(host_name string, event_handler_command string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_HOST_EVENT_HANDLER")
		Hostname(c, host_name)
		EventHandlerCommand(c, event_handler_command)
	}
}

// ChangeHostModattr is generated from the nagios external command definition:
// Desfinition: CHANGE_HOST_MODATTR;<host_name>;<value>
// Description:
//  This command changes the modified attributes value for the specified host.  Modified attributes values are used by Nagios to determine which object properties should be retained across program restarts.  Thus, modifying the value of the attributes can affect data retention.  This is an advanced option and should only be used by people who are intimately familiar with the data retention logic in Nagios.
func ChangeHostModattr(host_name string, value string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_HOST_MODATTR")
		Hostname(c, host_name)
		Value(c, value)
	}
}

// ChangeMaxHostCheckAttempts is generated from the nagios external command definition:
// Desfinition: CHANGE_MAX_HOST_CHECK_ATTEMPTS;<host_name>;<check_attempts>
// Description:
//  Changes the maximum number of check attempts (retries) for a particular host.
func ChangeMaxHostCheckAttempts(host_name string, check_attempts int) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_MAX_HOST_CHECK_ATTEMPTS")
		Hostname(c, host_name)
		CheckAttempts(c, check_attempts)
	}
}

// ChangeMaxSvcCheckAttempts is generated from the nagios external command definition:
// Desfinition: CHANGE_MAX_SVC_CHECK_ATTEMPTS;<host_name>;<service_description>;<check_attempts>
// Description:
//  Changes the maximum number of check attempts (retries) for a particular service.
func ChangeMaxSvcCheckAttempts(host_name string, service_description string, check_attempts int) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_MAX_SVC_CHECK_ATTEMPTS")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		CheckAttempts(c, check_attempts)
	}
}

// ChangeNormalHostCheckInterval is generated from the nagios external command definition:
// Desfinition: CHANGE_NORMAL_HOST_CHECK_INTERVAL;<host_name>;<check_interval>
// Description:
//  Changes the normal (regularly scheduled) check interval for a particular host.
func ChangeNormalHostCheckInterval(host_name string, check_interval time.Duration) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_NORMAL_HOST_CHECK_INTERVAL")
		Hostname(c, host_name)
		Duration(c, check_interval)
	}
}

// ChangeNormalSvcCheckInterval is generated from the nagios external command definition:
// Desfinition: CHANGE_NORMAL_SVC_CHECK_INTERVAL;<host_name>;<service_description>;<check_interval>
// Description:
//  Changes the normal (regularly scheduled) check interval for a particular service
func ChangeNormalSvcCheckInterval(host_name string, service_description string, check_interval time.Duration) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_NORMAL_SVC_CHECK_INTERVAL")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		Duration(c, check_interval)
	}
}

// ChangeRetryHostCheckInterval is generated from the nagios external command definition:
// Desfinition: CHANGE_RETRY_HOST_CHECK_INTERVAL;<host_name>;<service_description>;<check_interval>
// Description:
//  Changes the retry check interval for a particular host.
func ChangeRetryHostCheckInterval(host_name string, service_description string, check_interval time.Duration) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_RETRY_HOST_CHECK_INTERVAL")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		Duration(c, check_interval)
	}
}

// ChangeRetrySvcCheckInterval is generated from the nagios external command definition:
// Desfinition: CHANGE_RETRY_SVC_CHECK_INTERVAL;<host_name>;<service_description>;<check_interval>
// Description:
//  Changes the retry check interval for a particular service.
func ChangeRetrySvcCheckInterval(host_name string, service_description string, check_interval time.Duration) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_RETRY_SVC_CHECK_INTERVAL")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		Duration(c, check_interval)
	}
}

// ChangeSvcCheckCommand is generated from the nagios external command definition:
// Desfinition: CHANGE_SVC_CHECK_COMMAND;<host_name>;<service_description>;<check_command>
// Description:
//  Changes the check command for a particular service to be that specified by the "check_command" option.  The "check_command" option specifies the short name of the command that should be used as the new service check command.  The command must have been configured in Nagios before it was last (re)started.
func ChangeSvcCheckCommand(host_name string, service_description string, check_command string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_SVC_CHECK_COMMAND")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		CheckCommand(c, check_command)
	}
}

// ChangeSvcCheckTimeperiod is generated from the nagios external command definition:
// Desfinition: CHANGE_SVC_CHECK_TIMEPERIOD;<host_name>;<service_description>;<check_timeperiod>
// Description:
//  Changes the check timeperiod for a particular service to what is specified by the "check_timeperiod" option.  The "check_timeperiod" option should be the short name of the timeperod that is to be used as the service check timeperiod.  The timeperiod must have been configured in Nagios before it was last (re)started.
func ChangeSvcCheckTimeperiod(host_name string, service_description string, check_timeperiod string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_SVC_CHECK_TIMEPERIOD")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		CheckTimePeriod(c, check_timeperiod)
	}
}

// ChangeSvcEventHandler is generated from the nagios external command definition:
// Desfinition: CHANGE_SVC_EVENT_HANDLER;<host_name>;<service_description>;<event_handler_command>
// Description:
//  Changes the event handler command for a particular service to be that specified by the "event_handler_command" option.  The "event_handler_command" option specifies the short name of the command that should be used as the new service event handler.  The command must have been configured in Nagios before it was last (re)started.
func ChangeSvcEventHandler(host_name string, service_description string, event_handler_command string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_SVC_EVENT_HANDLER")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		EventHandlerCommand(c, event_handler_command)
	}
}

// ChangeSvcModattr is generated from the nagios external command definition:
// Desfinition: CHANGE_SVC_MODATTR;<host_name>;<service_description>;<value>
// Description:
//  This command changes the modified attributes value for the specified service.  Modified attributes values are used by Nagios to determine which object properties should be retained across program restarts.  Thus, modifying the value of the attributes can affect data retention.  This is an advanced option and should only be used by people who are intimately familiar with the data retention logic in Nagios.
func ChangeSvcModattr(host_name string, service_description string, value string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_SVC_MODATTR")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		Value(c, value)
	}
}

// ChangeSvcNotificationTimeperiod is generated from the nagios external command definition:
// Desfinition: CHANGE_SVC_NOTIFICATION_TIMEPERIOD;<host_name>;<service_description>;<notification_timeperiod>
// Description:
//  Changes the notification timeperiod for a particular service to what is specified by the "notification_timeperiod" option.  The "notification_timeperiod" option should be the short name of the timeperiod that is to be used as the service notification timeperiod.  The timeperiod must have been configured in Nagios before it was last (re)started.
func ChangeSvcNotificationTimeperiod(host_name string, service_description string, notification_timeperiod string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("CHANGE_SVC_NOTIFICATION_TIMEPERIOD")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		NotificationTimePeriod(c, notification_timeperiod)
	}
}

// DelayHostNotification is generated from the nagios external command definition:
// Desfinition: DELAY_HOST_NOTIFICATION;<host_name>;<notification_time>
// Description:
//  Delays the next notification for a parciular service until "notification_time".  The "notification_time" argument is specified in time_t format (seconds since the UNIX epoch).  Note that this will only have an affect if the service stays in the same problem state that it is currently in.  If the service changes to another state, a new notification may go out before the time you specify in the "notification_time" argument.
func DelayHostNotification(host_name string, notification_time time.Time) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DELAY_HOST_NOTIFICATION")
		Hostname(c, host_name)
		NotificationTime(c, notification_time)
	}
}

// DelaySvcNotification is generated from the nagios external command definition:
// Desfinition: DELAY_SVC_NOTIFICATION;<host_name>;<service_description>;<notification_time>
// Description:
//  Delays the next notification for a parciular service until "notification_time".  The "notification_time" argument is specified in time_t format (seconds since the UNIX epoch).  Note that this will only have an affect if the service stays in the same problem state that it is currently in.  If the service changes to another state, a new notification may go out before the time you specify in the "notification_time" argument.
func DelaySvcNotification(host_name string, service_description string, notification_time time.Time) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DELAY_SVC_NOTIFICATION")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		NotificationTime(c, notification_time)
	}
}

// DelAllHostComments is generated from the nagios external command definition:
// Desfinition: DEL_ALL_HOST_COMMENTS;<host_name>
// Description:
//  Deletes all comments assocated with a particular host.
func DelAllHostComments(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DEL_ALL_HOST_COMMENTS")
		Hostname(c, host_name)
	}
}

// DelAllSvcComments is generated from the nagios external command definition:
// Desfinition: DEL_ALL_SVC_COMMENTS;<host_name>;<service_description>
// Description:
//  Deletes all comments associated with a particular service.
func DelAllSvcComments(host_name string, service_description string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DEL_ALL_SVC_COMMENTS")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
	}
}

// DelHostComment is generated from the nagios external command definition:
// Desfinition: DEL_HOST_COMMENT;<comment_id>
// Description:
//  Deletes a host comment.  The id number of the comment that is to be deleted must be specified.
func DelHostComment(comment_id int) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DEL_HOST_COMMENT")
		CommentID(c, comment_id)
	}
}

// DelHostDowntime is generated from the nagios external command definition:
// Desfinition: DEL_HOST_DOWNTIME;<downtime_id>
// Description:
//  Deletes the host downtime entry that has an ID number matching the "downtime_id" argument.  If the downtime is currently in effect, the host will come out of scheduled downtime (as long as there are no other overlapping active downtime entries).
func DelHostDowntime(downtime_id int) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DEL_HOST_DOWNTIME")
		DowntimeID(c, downtime_id)
	}
}

// DelSvcComment is generated from the nagios external command definition:
// Desfinition: DEL_SVC_COMMENT;<comment_id>
// Description:
//  Deletes a service comment.  The id number of the comment that is to be deleted must be specified.
func DelSvcComment(comment_id int) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DEL_SVC_COMMENT")
		CommentID(c, comment_id)
	}
}

// DelSvcDowntime is generated from the nagios external command definition:
// Desfinition: DEL_SVC_DOWNTIME;<downtime_id>
// Description:
//  Deletes the service downtime entry that has an ID number matching the "downtime_id" argument.  If the downtime is currently in effect, the service will come out of scheduled downtime (as long as there are no other overlapping active downtime entries).
func DelSvcDowntime(downtime_id int) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DEL_SVC_DOWNTIME")
		DowntimeID(c, downtime_id)
	}
}

// DisableAllNotificationsBeyondHost is generated from the nagios external command definition:
// Desfinition: DISABLE_ALL_NOTIFICATIONS_BEYOND_HOST;<host_name>
// Description:
//  Disables notifications for all hosts and services "beyond" (e.g. on all child hosts of) the specified host.  The current notification setting for the specified host is not affected.
func DisableAllNotificationsBeyondHost(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_ALL_NOTIFICATIONS_BEYOND_HOST")
		Hostname(c, host_name)
	}
}

// DisableContactgroupHostNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_CONTACTGROUP_HOST_NOTIFICATIONS;<contactgroup_name>
// Description:
//  Disables host notifications for all contacts in a particular contactgroup.
func DisableContactgroupHostNotifications(contactgroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_CONTACTGROUP_HOST_NOTIFICATIONS")
		ContactGroupName(c, contactgroup_name)
	}
}

// DisableContactgroupSvcNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_CONTACTGROUP_SVC_NOTIFICATIONS;<contactgroup_name>
// Description:
//  Disables service notifications for all contacts in a particular contactgroup.
func DisableContactgroupSvcNotifications(contactgroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_CONTACTGROUP_SVC_NOTIFICATIONS")
		ContactGroupName(c, contactgroup_name)
	}
}

// DisableContactHostNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_CONTACT_HOST_NOTIFICATIONS;<contact_name>
// Description:
//  Disables host notifications for a particular contact.
func DisableContactHostNotifications(contact_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_CONTACT_HOST_NOTIFICATIONS")
		ContactName(c, contact_name)
	}
}

// DisableContactSvcNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_CONTACT_SVC_NOTIFICATIONS;<contact_name>
// Description:
//  Disables service notifications for a particular contact.
func DisableContactSvcNotifications(contact_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_CONTACT_SVC_NOTIFICATIONS")
		ContactName(c, contact_name)
	}
}

// DisableEventHandlers is generated from the nagios external command definition:
// Desfinition: DISABLE_EVENT_HANDLERS
// Description:
//  Disables host and service event handlers on a program-wide basis.
func DisableEventHandlers() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_EVENT_HANDLERS")
	}
}

// DisableFailurePrediction is generated from the nagios external command definition:
// Desfinition: DISABLE_FAILURE_PREDICTION
// Description:
//  Disables failure prediction on a program-wide basis.  This feature is not currently implemented in Nagios.
func DisableFailurePrediction() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_FAILURE_PREDICTION")
	}
}

// DisableFlapDetection is generated from the nagios external command definition:
// Desfinition: DISABLE_FLAP_DETECTION
// Description:
//  Disables host and service flap detection on a program-wide basis.
func DisableFlapDetection() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_FLAP_DETECTION")
	}
}

// DisableHostgroupHostChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_HOSTGROUP_HOST_CHECKS;<hostgroup_name>
// Description:
//  Disables active checks for all hosts in a particular hostgroup.
func DisableHostgroupHostChecks(hostgroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_HOSTGROUP_HOST_CHECKS")
		HostGroupName(c, hostgroup_name)
	}
}

// DisableHostgroupHostNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_HOSTGROUP_HOST_NOTIFICATIONS;<hostgroup_name>
// Description:
//  Disables notifications for all hosts in a particular hostgroup.  This does not disable notifications for the services associated with the hosts in the hostgroup - see the DISABLE_HOSTGROUP_SVC_NOTIFICATIONS command for that.
func DisableHostgroupHostNotifications(hostgroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_HOSTGROUP_HOST_NOTIFICATIONS")
		HostGroupName(c, hostgroup_name)
	}
}

// DisableHostgroupPassiveHostChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_HOSTGROUP_PASSIVE_HOST_CHECKS;<hostgroup_name>
// Description:
//  Disables passive checks for all hosts in a particular hostgroup.
func DisableHostgroupPassiveHostChecks(hostgroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_HOSTGROUP_PASSIVE_HOST_CHECKS")
		HostGroupName(c, hostgroup_name)
	}
}

// DisableHostgroupPassiveSvcChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_HOSTGROUP_PASSIVE_SVC_CHECKS;<hostgroup_name>
// Description:
//  Disables passive checks for all services associated with hosts in a particular hostgroup.
func DisableHostgroupPassiveSvcChecks(hostgroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_HOSTGROUP_PASSIVE_SVC_CHECKS")
		HostGroupName(c, hostgroup_name)
	}
}

// DisableHostgroupSvcChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_HOSTGROUP_SVC_CHECKS;<hostgroup_name>
// Description:
//  Disables active checks for all services associated with hosts in a particular hostgroup.
func DisableHostgroupSvcChecks(hostgroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_HOSTGROUP_SVC_CHECKS")
		HostGroupName(c, hostgroup_name)
	}
}

// DisableHostgroupSvcNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_HOSTGROUP_SVC_NOTIFICATIONS;<hostgroup_name>
// Description:
//  Disables notifications for all services associated with hosts in a particular hostgroup.  This does not disable notifications for the hosts in the hostgroup - see the DISABLE_HOSTGROUP_HOST_NOTIFICATIONS command for that.
func DisableHostgroupSvcNotifications(hostgroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_HOSTGROUP_SVC_NOTIFICATIONS")
		HostGroupName(c, hostgroup_name)
	}
}

// DisableHostAndChildNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_HOST_AND_CHILD_NOTIFICATIONS;<host_name>
// Description:
//  Disables notifications for the specified host, as well as all hosts "beyond" (e.g. on all child hosts of) the specified host.
func DisableHostAndChildNotifications(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_HOST_AND_CHILD_NOTIFICATIONS")
		Hostname(c, host_name)
	}
}

// DisableHostCheck is generated from the nagios external command definition:
// Desfinition: DISABLE_HOST_CHECK;<host_name>
// Description:
//  Disables (regularly scheduled and on-demand) active checks of the specified host.
func DisableHostCheck(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_HOST_CHECK")
		Hostname(c, host_name)
	}
}

// DisableHostEventHandler is generated from the nagios external command definition:
// Desfinition: DISABLE_HOST_EVENT_HANDLER;<host_name>
// Description:
//  Disables the event handler for the specified host.
func DisableHostEventHandler(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_HOST_EVENT_HANDLER")
		Hostname(c, host_name)
	}
}

// DisableHostFlapDetection is generated from the nagios external command definition:
// Desfinition: DISABLE_HOST_FLAP_DETECTION;<host_name>
// Description:
//  Disables flap detection for the specified host.
func DisableHostFlapDetection(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_HOST_FLAP_DETECTION")
		Hostname(c, host_name)
	}
}

// DisableHostFreshnessChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_HOST_FRESHNESS_CHECKS
// Description:
//  Disables freshness checks of all hosts on a program-wide basis.
func DisableHostFreshnessChecks() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_HOST_FRESHNESS_CHECKS")
	}
}

// DisableHostNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_HOST_NOTIFICATIONS;<host_name>
// Description:
//  Disables notifications for a particular host.
func DisableHostNotifications(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_HOST_NOTIFICATIONS")
		Hostname(c, host_name)
	}
}

// DisableHostSvcChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_HOST_SVC_CHECKS;<host_name>
// Description:
//  Enables active checks of all services on the specified host.
func DisableHostSvcChecks(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_HOST_SVC_CHECKS")
		Hostname(c, host_name)
	}
}

// DisableHostSvcNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_HOST_SVC_NOTIFICATIONS;<host_name>
// Description:
//  Disables notifications for all services on the specified host.
func DisableHostSvcNotifications(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_HOST_SVC_NOTIFICATIONS")
		Hostname(c, host_name)
	}
}

// DisableNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_NOTIFICATIONS
// Description:
//  Disables host and service notifications on a program-wide basis.
func DisableNotifications() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_NOTIFICATIONS")
	}
}

// DisablePassiveHostChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_PASSIVE_HOST_CHECKS;<host_name>
// Description:
//  Disables acceptance and processing of passive host checks for the specified host.
func DisablePassiveHostChecks(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_PASSIVE_HOST_CHECKS")
		Hostname(c, host_name)
	}
}

// DisablePassiveSvcChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_PASSIVE_SVC_CHECKS;<host_name>;<service_description>
// Description:
//  Disables passive checks for the specified service.
func DisablePassiveSvcChecks(host_name string, service_description string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_PASSIVE_SVC_CHECKS")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
	}
}

// DisablePerformanceData is generated from the nagios external command definition:
// Desfinition: DISABLE_PERFORMANCE_DATA
// Description:
//  Disables the processing of host and service performance data on a program-wide basis.
func DisablePerformanceData() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_PERFORMANCE_DATA")
	}
}

// DisableServicegroupHostChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_SERVICEGROUP_HOST_CHECKS;<servicegroup_name>
// Description:
//  Disables active checks for all hosts that have services that are members of a particular hostgroup.
func DisableServicegroupHostChecks(servicegroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_SERVICEGROUP_HOST_CHECKS")
		ServiceGroupName(c, servicegroup_name)
	}
}

// DisableServicegroupHostNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_SERVICEGROUP_HOST_NOTIFICATIONS;<servicegroup_name>
// Description:
//  Disables notifications for all hosts that have services that are members of a particular servicegroup.
func DisableServicegroupHostNotifications(servicegroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_SERVICEGROUP_HOST_NOTIFICATIONS")
		ServiceGroupName(c, servicegroup_name)
	}
}

// DisableServicegroupPassiveHostChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_SERVICEGROUP_PASSIVE_HOST_CHECKS;<servicegroup_name>
// Description:
//  Disables the acceptance and processing of passive checks for all hosts that have services that are members of a particular service group.
func DisableServicegroupPassiveHostChecks(servicegroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_SERVICEGROUP_PASSIVE_HOST_CHECKS")
		ServiceGroupName(c, servicegroup_name)
	}
}

// DisableServicegroupPassiveSvcChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_SERVICEGROUP_PASSIVE_SVC_CHECKS;<servicegroup_name>
// Description:
//  Disables the acceptance and processing of passive checks for all services in a particular servicegroup.
func DisableServicegroupPassiveSvcChecks(servicegroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_SERVICEGROUP_PASSIVE_SVC_CHECKS")
		ServiceGroupName(c, servicegroup_name)
	}
}

// DisableServicegroupSvcChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_SERVICEGROUP_SVC_CHECKS;<servicegroup_name>
// Description:
//  Disables active checks for all services in a particular servicegroup.
func DisableServicegroupSvcChecks(servicegroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_SERVICEGROUP_SVC_CHECKS")
		ServiceGroupName(c, servicegroup_name)
	}
}

// DisableServicegroupSvcNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_SERVICEGROUP_SVC_NOTIFICATIONS;<servicegroup_name>
// Description:
//  Disables notifications for all services that are members of a particular servicegroup.
func DisableServicegroupSvcNotifications(servicegroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_SERVICEGROUP_SVC_NOTIFICATIONS")
		ServiceGroupName(c, servicegroup_name)
	}
}

// DisableServiceFlapDetection is generated from the nagios external command definition:
// Desfinition: DISABLE_SERVICE_FLAP_DETECTION;<host_name>;<service_description>
// Description:
//  Disables flap detection for the specified service.
func DisableServiceFlapDetection(host_name string, service_description string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_SERVICE_FLAP_DETECTION")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
	}
}

// DisableServiceFreshnessChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_SERVICE_FRESHNESS_CHECKS
// Description:
//  Disables freshness checks of all services on a program-wide basis.
func DisableServiceFreshnessChecks() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_SERVICE_FRESHNESS_CHECKS")
	}
}

// DisableSvcCheck is generated from the nagios external command definition:
// Desfinition: DISABLE_SVC_CHECK;<host_name>;<service_description>
// Description:
//  Disables active checks for a particular service.
func DisableSvcCheck(host_name string, service_description string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_SVC_CHECK")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
	}
}

// DisableSvcEventHandler is generated from the nagios external command definition:
// Desfinition: DISABLE_SVC_EVENT_HANDLER;<host_name>;<service_description>
// Description:
//  Disables the event handler for the specified service.
func DisableSvcEventHandler(host_name string, service_description string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_SVC_EVENT_HANDLER")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
	}
}

// DisableSvcFlapDetection is generated from the nagios external command definition:
// Desfinition: DISABLE_SVC_FLAP_DETECTION;<host_name>;<service_description>
// Description:
//  Disables flap detection for the specified service.
func DisableSvcFlapDetection(host_name string, service_description string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_SVC_FLAP_DETECTION")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
	}
}

// DisableSvcNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_SVC_NOTIFICATIONS;<host_name>;<service_description>
// Description:
//  Disables notifications for a particular service.
func DisableSvcNotifications(host_name string, service_description string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("DISABLE_SVC_NOTIFICATIONS")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
	}
}

// EnableAllNotificationsBeyondHost is generated from the nagios external command definition:
// Desfinition: ENABLE_ALL_NOTIFICATIONS_BEYOND_HOST;<host_name>
// Description:
//  Enables notifications for all hosts and services "beyond" (e.g. on all child hosts of) the specified host.  The current notification setting for the specified host is not affected.  Notifications will only be sent out for these hosts and services if notifications are also enabled on a program-wide basis.
func EnableAllNotificationsBeyondHost(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_ALL_NOTIFICATIONS_BEYOND_HOST")
		Hostname(c, host_name)
	}
}

// EnableContactgroupHostNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_CONTACTGROUP_HOST_NOTIFICATIONS;<contactgroup_name>
// Description:
//  Enables host notifications for all contacts in a particular contactgroup.
func EnableContactgroupHostNotifications(contactgroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_CONTACTGROUP_HOST_NOTIFICATIONS")
		ContactGroupName(c, contactgroup_name)
	}
}

// EnableContactgroupSvcNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_CONTACTGROUP_SVC_NOTIFICATIONS;<contactgroup_name>
// Description:
//  Enables service notifications for all contacts in a particular contactgroup.
func EnableContactgroupSvcNotifications(contactgroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_CONTACTGROUP_SVC_NOTIFICATIONS")
		ContactGroupName(c, contactgroup_name)
	}
}

// EnableContactHostNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_CONTACT_HOST_NOTIFICATIONS;<contact_name>
// Description:
//  Enables host notifications for a particular contact.
func EnableContactHostNotifications(contact_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_CONTACT_HOST_NOTIFICATIONS")
		ContactName(c, contact_name)
	}
}

// EnableContactSvcNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_CONTACT_SVC_NOTIFICATIONS;<contact_name>
// Description:
//  Disables service notifications for a particular contact.
func EnableContactSvcNotifications(contact_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_CONTACT_SVC_NOTIFICATIONS")
		ContactName(c, contact_name)
	}
}

// EnableEventHandlers is generated from the nagios external command definition:
// Desfinition: ENABLE_EVENT_HANDLERS
// Description:
//  Enables host and service event handlers on a program-wide basis.
func EnableEventHandlers() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_EVENT_HANDLERS")
	}
}

// EnableFailurePrediction is generated from the nagios external command definition:
// Desfinition: ENABLE_FAILURE_PREDICTION
// Description:
//  Enables failure prediction on a program-wide basis.  This feature is not currently implemented in Nagios.
func EnableFailurePrediction() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_FAILURE_PREDICTION")
	}
}

// EnableFlapDetection is generated from the nagios external command definition:
// Desfinition: ENABLE_FLAP_DETECTION
// Description:
//  Enables host and service flap detection on a program-wide basis.
func EnableFlapDetection() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_FLAP_DETECTION")
	}
}

// EnableHostgroupHostChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_HOSTGROUP_HOST_CHECKS;<hostgroup_name>
// Description:
//  Enables active checks for all hosts in a particular hostgroup.
func EnableHostgroupHostChecks(hostgroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_HOSTGROUP_HOST_CHECKS")
		HostGroupName(c, hostgroup_name)
	}
}

// EnableHostgroupHostNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_HOSTGROUP_HOST_NOTIFICATIONS;<hostgroup_name>
// Description:
//  Enables notifications for all hosts in a particular hostgroup.  This does not enable notifications for the services associated with the hosts in the hostgroup - see the ENABLE_HOSTGROUP_SVC_NOTIFICATIONS command for that.  In order for notifications to be sent out for these hosts, notifications must be enabled on a program-wide basis as well.
func EnableHostgroupHostNotifications(hostgroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_HOSTGROUP_HOST_NOTIFICATIONS")
		HostGroupName(c, hostgroup_name)
	}
}

// EnableHostgroupPassiveHostChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_HOSTGROUP_PASSIVE_HOST_CHECKS;<hostgroup_name>
// Description:
//  Enables passive checks for all hosts in a particular hostgroup.
func EnableHostgroupPassiveHostChecks(hostgroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_HOSTGROUP_PASSIVE_HOST_CHECKS")
		HostGroupName(c, hostgroup_name)
	}
}

// EnableHostgroupPassiveSvcChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_HOSTGROUP_PASSIVE_SVC_CHECKS;<hostgroup_name>
// Description:
//  Enables passive checks for all services associated with hosts in a particular hostgroup.
func EnableHostgroupPassiveSvcChecks(hostgroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_HOSTGROUP_PASSIVE_SVC_CHECKS")
		HostGroupName(c, hostgroup_name)
	}
}

// EnableHostgroupSvcChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_HOSTGROUP_SVC_CHECKS;<hostgroup_name>
// Description:
//  Enables active checks for all services associated with hosts in a particular hostgroup.
func EnableHostgroupSvcChecks(hostgroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_HOSTGROUP_SVC_CHECKS")
		HostGroupName(c, hostgroup_name)
	}
}

// EnableHostgroupSvcNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_HOSTGROUP_SVC_NOTIFICATIONS;<hostgroup_name>
// Description:
//  Enables notifications for all services that are associated with hosts in a particular hostgroup.  This does not enable notifications for the hosts in the hostgroup - see the ENABLE_HOSTGROUP_HOST_NOTIFICATIONS command for that.  In order for notifications to be sent out for these services, notifications must be enabled on a program-wide basis as well.
func EnableHostgroupSvcNotifications(hostgroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_HOSTGROUP_SVC_NOTIFICATIONS")
		HostGroupName(c, hostgroup_name)
	}
}

// EnableHostAndChildNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_HOST_AND_CHILD_NOTIFICATIONS;<host_name>
// Description:
//  Enables notifications for the specified host, as well as all hosts "beyond" (e.g. on all child hosts of) the specified host.  Notifications will only be sent out for these hosts if notifications are also enabled on a program-wide basis.
func EnableHostAndChildNotifications(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_HOST_AND_CHILD_NOTIFICATIONS")
		Hostname(c, host_name)
	}
}

// EnableHostCheck is generated from the nagios external command definition:
// Desfinition: ENABLE_HOST_CHECK;<host_name>
// Description:
//  Enables (regularly scheduled and on-demand) active checks of the specified host.
func EnableHostCheck(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_HOST_CHECK")
		Hostname(c, host_name)
	}
}

// EnableHostEventHandler is generated from the nagios external command definition:
// Desfinition: ENABLE_HOST_EVENT_HANDLER;<host_name>
// Description:
//  Enables the event handler for the specified host.
func EnableHostEventHandler(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_HOST_EVENT_HANDLER")
		Hostname(c, host_name)
	}
}

// EnableHostFlapDetection is generated from the nagios external command definition:
// Desfinition: ENABLE_HOST_FLAP_DETECTION;<host_name>
// Description:
//  Enables flap detection for the specified host.  In order for the flap detection algorithms to be run for the host, flap detection must be enabled on a program-wide basis as well.
func EnableHostFlapDetection(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_HOST_FLAP_DETECTION")
		Hostname(c, host_name)
	}
}

// EnableHostFreshnessChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_HOST_FRESHNESS_CHECKS
// Description:
//  Enables freshness checks of all hosts on a program-wide basis.  Individual hosts that have freshness checks disabled will not be checked for freshness.
func EnableHostFreshnessChecks() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_HOST_FRESHNESS_CHECKS")
	}
}

// EnableHostNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_HOST_NOTIFICATIONS;<host_name>
// Description:
//  Enables notifications for a particular host.  Notifications will be sent out for the host only if notifications are enabled on a program-wide basis as well.
func EnableHostNotifications(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_HOST_NOTIFICATIONS")
		Hostname(c, host_name)
	}
}

// EnableHostSvcChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_HOST_SVC_CHECKS;<host_name>
// Description:
//  Enables active checks of all services on the specified host.
func EnableHostSvcChecks(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_HOST_SVC_CHECKS")
		Hostname(c, host_name)
	}
}

// EnableHostSvcNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_HOST_SVC_NOTIFICATIONS;<host_name>
// Description:
//  Enables notifications for all services on the specified host.  Note that notifications will not be sent out if notifications are disabled on a program-wide basis.
func EnableHostSvcNotifications(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_HOST_SVC_NOTIFICATIONS")
		Hostname(c, host_name)
	}
}

// EnableNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_NOTIFICATIONS
// Description:
//  Enables host and service notifications on a program-wide basis.
func EnableNotifications() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_NOTIFICATIONS")
	}
}

// EnablePassiveHostChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_PASSIVE_HOST_CHECKS;<host_name>
// Description:
//  Enables acceptance and processing of passive host checks for the specified host.
func EnablePassiveHostChecks(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_PASSIVE_HOST_CHECKS")
		Hostname(c, host_name)
	}
}

// EnablePassiveSvcChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_PASSIVE_SVC_CHECKS;<host_name>;<service_description>
// Description:
//  Enables passive checks for the specified service.
func EnablePassiveSvcChecks(host_name string, service_description string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_PASSIVE_SVC_CHECKS")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
	}
}

// EnablePerformanceData is generated from the nagios external command definition:
// Desfinition: ENABLE_PERFORMANCE_DATA
// Description:
//  Enables the processing of host and service performance data on a program-wide basis.
func EnablePerformanceData() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_PERFORMANCE_DATA")
	}
}

// EnableServicegroupHostChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_SERVICEGROUP_HOST_CHECKS;<servicegroup_name>
// Description:
//  Enables active checks for all hosts that have services that are members of a particular hostgroup.
func EnableServicegroupHostChecks(servicegroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_SERVICEGROUP_HOST_CHECKS")
		ServiceGroupName(c, servicegroup_name)
	}
}

// EnableServicegroupHostNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_SERVICEGROUP_HOST_NOTIFICATIONS;<servicegroup_name>
// Description:
//  Enables notifications for all hosts that have services that are members of a particular servicegroup.  In order for notifications to be sent out for these hosts, notifications must also be enabled on a program-wide basis.
func EnableServicegroupHostNotifications(servicegroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_SERVICEGROUP_HOST_NOTIFICATIONS")
		ServiceGroupName(c, servicegroup_name)
	}
}

// EnableServicegroupPassiveHostChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_SERVICEGROUP_PASSIVE_HOST_CHECKS;<servicegroup_name>
// Description:
//  Enables the acceptance and processing of passive checks for all hosts that have services that are members of a particular service group.
func EnableServicegroupPassiveHostChecks(servicegroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_SERVICEGROUP_PASSIVE_HOST_CHECKS")
		ServiceGroupName(c, servicegroup_name)
	}
}

// EnableServicegroupPassiveSvcChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_SERVICEGROUP_PASSIVE_SVC_CHECKS;<servicegroup_name>
// Description:
//  Enables the acceptance and processing of passive checks for all services in a particular servicegroup.
func EnableServicegroupPassiveSvcChecks(servicegroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_SERVICEGROUP_PASSIVE_SVC_CHECKS")
		ServiceGroupName(c, servicegroup_name)
	}
}

// EnableServicegroupSvcChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_SERVICEGROUP_SVC_CHECKS;<servicegroup_name>
// Description:
//  Enables active checks for all services in a particular servicegroup.
func EnableServicegroupSvcChecks(servicegroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_SERVICEGROUP_SVC_CHECKS")
		ServiceGroupName(c, servicegroup_name)
	}
}

// EnableServicegroupSvcNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_SERVICEGROUP_SVC_NOTIFICATIONS;<servicegroup_name>
// Description:
//  Enables notifications for all services that are members of a particular servicegroup.  In order for notifications to be sent out for these services, notifications must also be enabled on a program-wide basis.
func EnableServicegroupSvcNotifications(servicegroup_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_SERVICEGROUP_SVC_NOTIFICATIONS")
		ServiceGroupName(c, servicegroup_name)
	}
}

// EnableServiceFreshnessChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_SERVICE_FRESHNESS_CHECKS
// Description:
//  Enables freshness checks of all services on a program-wide basis.  Individual services that have freshness checks disabled will not be checked for freshness.
func EnableServiceFreshnessChecks() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_SERVICE_FRESHNESS_CHECKS")
	}
}

// EnableSvcCheck is generated from the nagios external command definition:
// Desfinition: ENABLE_SVC_CHECK;<host_name>;<service_description>
// Description:
//  Enables active checks for a particular service.
func EnableSvcCheck(host_name string, service_description string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_SVC_CHECK")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
	}
}

// EnableSvcEventHandler is generated from the nagios external command definition:
// Desfinition: ENABLE_SVC_EVENT_HANDLER;<host_name>;<service_description>
// Description:
//  Enables the event handler for the specified service.
func EnableSvcEventHandler(host_name string, service_description string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_SVC_EVENT_HANDLER")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
	}
}

// EnableSvcFlapDetection is generated from the nagios external command definition:
// Desfinition: ENABLE_SVC_FLAP_DETECTION;<host_name>;<service_description>
// Description:
//  Enables flap detection for the specified service.  In order for the flap detection algorithms to be run for the service, flap detection must be enabled on a program-wide basis as well.
func EnableSvcFlapDetection(host_name string, service_description string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_SVC_FLAP_DETECTION")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
	}
}

// EnableSvcNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_SVC_NOTIFICATIONS;<host_name>;<service_description>
// Description:
//  Enables notifications for a particular service.  Notifications will be sent out for the service only if notifications are enabled on a program-wide basis as well.
func EnableSvcNotifications(host_name string, service_description string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("ENABLE_SVC_NOTIFICATIONS")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
	}
}

// ProcessFile is generated from the nagios external command definition:
// Desfinition: PROCESS_FILE;<file_name>;<delete>
// Description:
//  Directs Nagios to process all external commands that are found in the file specified by the <file_name> argument.  If the <delete> option is non-zero, the file will be deleted once it has been processes.  If the <delete> option is set to zero, the file is left untouched.
func ProcessFile(file_name string, delete bool) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("PROCESS_FILE")
		FileName(c, file_name)
		Delete(c, delete)
	}
}

// ProcessHostCheckResult is generated from the nagios external command definition:
// Desfinition: PROCESS_HOST_CHECK_RESULT;<host_name>;<status_code>;<plugin_output>
// Description:
//  This is used to submit a passive check result for a particular host.  The "status_code" indicates the state of the host check and should be one of the following: 0=UP, 1=DOWN, 2=UNREACHABLE.  The "plugin_output" argument contains the text returned from the host check, along with optional performance data.
func ProcessHostCheckResult(host_name string, status_code int, plugin_output string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("PROCESS_HOST_CHECK_RESULT")
		Hostname(c, host_name)
		StatusCode(c, status_code)
		PluginOutput(c, plugin_output)
	}
}

// ProcessServiceCheckResult is generated from the nagios external command definition:
// Desfinition: PROCESS_SERVICE_CHECK_RESULT;<host_name>;<service_description>;<return_code>;<plugin_output>
// Description:
//  This is used to submit a passive check result for a particular service.  The "return_code" field should be one of the following: 0=OK, 1=WARNING, 2=CRITICAL, 3=UNKNOWN.  The "plugin_output" field contains text output from the service check, along with optional performance data.
func ProcessServiceCheckResult(host_name string, service_description string, return_code int, plugin_output string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("PROCESS_SERVICE_CHECK_RESULT")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		ReturnCode(c, return_code)
		PluginOutput(c, plugin_output)
	}
}

// ReadStateInformation is generated from the nagios external command definition:
// Desfinition: READ_STATE_INFORMATION
// Description:
//  Causes Nagios to load all current monitoring status information from the state retention file.  Normally, state retention information is loaded when the Nagios process starts up and before it starts monitoring.  WARNING: This command will cause Nagios to discard all current monitoring status information and use the information stored in state retention file!  Use with care.
func ReadStateInformation() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("READ_STATE_INFORMATION")
	}
}

// RemoveHostAcknowledgement is generated from the nagios external command definition:
// Desfinition: REMOVE_HOST_ACKNOWLEDGEMENT;<host_name>
// Description:
//  This removes the problem acknowledgement for a particular host.  Once the acknowledgement has been removed, notifications can once again be sent out for the given host.
func RemoveHostAcknowledgement(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("REMOVE_HOST_ACKNOWLEDGEMENT")
		Hostname(c, host_name)
	}
}

// RemoveSvcAcknowledgement is generated from the nagios external command definition:
// Desfinition: REMOVE_SVC_ACKNOWLEDGEMENT;<host_name>;<service_description>
// Description:
//  This removes the problem acknowledgement for a particular service.  Once the acknowledgement has been removed, notifications can once again be sent out for the given service.
func RemoveSvcAcknowledgement(host_name string, service_description string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("REMOVE_SVC_ACKNOWLEDGEMENT")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
	}
}

// RestartProgram is generated from the nagios external command definition:
// Desfinition: RESTART_PROGRAM
// Description:
//  Restarts the Nagios process.
func RestartProgram() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("RESTART_PROGRAM")
	}
}

// SaveStateInformation is generated from the nagios external command definition:
// Desfinition: SAVE_STATE_INFORMATION
// Description:
//  Causes Nagios to save all current monitoring status information to the state retention file.  Normally, state retention information is saved before the Nagios process shuts down and (potentially) at regularly scheduled intervals.  This command allows you to force Nagios to save this information to the state retention file immediately.  This does not affect the current status information in the Nagios process.
func SaveStateInformation() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SAVE_STATE_INFORMATION")
	}
}

// ScheduleAndPropagateHostDowntime is generated from the nagios external command definition:
// Desfinition: SCHEDULE_AND_PROPAGATE_HOST_DOWNTIME;<host_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
// Description:
//  Schedules downtime for a specified host and all of its children (hosts).  If the "fixed" argument is set to one (1), downtime will start and end at the times specified by the "start" and "end" arguments.  Otherwise, downtime will begin between the "start" and "end" times and last for "duration" seconds.  The "start" and "end" arguments are specified in time_t format (seconds since the UNIX epoch).  The specified (parent) host downtime can be triggered by another downtime entry if the "trigger_id" is set to the ID of another scheduled downtime entry.  Set the "trigger_id" argument to zero (0) if the downtime for the specified (parent) host should not be triggered by another downtime entry.
func ScheduleAndPropagateHostDowntime(host_name string, start_time time.Time, end_time time.Time, fixed bool, trigger_id int, duration time.Duration, author string, comment string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SCHEDULE_AND_PROPAGATE_HOST_DOWNTIME")
		Hostname(c, host_name)
		Start(c, start_time)
		End(c, end_time)
		Fixed(c, fixed)
		TriggerID(c, trigger_id)
		Duration(c, duration)
		Author(c, author)
		Comment(c, comment)
	}
}

// ScheduleAndPropagateTriggeredHostDowntime is generated from the nagios external command definition:
// Desfinition: SCHEDULE_AND_PROPAGATE_TRIGGERED_HOST_DOWNTIME;<host_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
// Description:
//  Schedules downtime for a specified host and all of its children (hosts).  If the "fixed" argument is set to one (1), downtime will start and end at the times specified by the "start" and "end" arguments.  Otherwise, downtime will begin between the "start" and "end" times and last for "duration" seconds.  The "start" and "end" arguments are specified in time_t format (seconds since the UNIX epoch).  Downtime for child hosts are all set to be triggered by the downtime for the specified (parent) host.  The specified (parent) host downtime can be triggered by another downtime entry if the "trigger_id" is set to the ID of another scheduled downtime entry.  Set the "trigger_id" argument to zero (0) if the downtime for the specified (parent) host should not be triggered by another downtime entry.
func ScheduleAndPropagateTriggeredHostDowntime(host_name string, start_time time.Time, end_time time.Time, fixed bool, trigger_id int, duration time.Duration, author string, comment string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SCHEDULE_AND_PROPAGATE_TRIGGERED_HOST_DOWNTIME")
		Hostname(c, host_name)
		Start(c, start_time)
		End(c, end_time)
		Fixed(c, fixed)
		TriggerID(c, trigger_id)
		Duration(c, duration)
		Author(c, author)
		Comment(c, comment)
	}
}

// ScheduleForcedHostCheck is generated from the nagios external command definition:
// Desfinition: SCHEDULE_FORCED_HOST_CHECK;<host_name>;<check_time>
// Description:
//  Schedules a forced active check of a particular host at "check_time".  The "check_time" argument is specified in time_t format (seconds since the UNIX epoch).   Forced checks are performed regardless of what time it is (e.g. timeperiod restrictions are ignored) and whether or not active checks are enabled on a host-specific or program-wide basis.
func ScheduleForcedHostCheck(host_name string, check_time time.Time) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SCHEDULE_FORCED_HOST_CHECK")
		Hostname(c, host_name)
		CheckTime(c, check_time)
	}
}

// ScheduleForcedHostSvcChecks is generated from the nagios external command definition:
// Desfinition: SCHEDULE_FORCED_HOST_SVC_CHECKS;<host_name>;<check_time>
// Description:
//  Schedules a forced active check of all services associated with a particular host at "check_time".  The "check_time" argument is specified in time_t format (seconds since the UNIX epoch).   Forced checks are performed regardless of what time it is (e.g. timeperiod restrictions are ignored) and whether or not active checks are enabled on a service-specific or program-wide basis.
func ScheduleForcedHostSvcChecks(host_name string, check_time time.Time) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SCHEDULE_FORCED_HOST_SVC_CHECKS")
		Hostname(c, host_name)
		CheckTime(c, check_time)
	}
}

// ScheduleForcedSvcCheck is generated from the nagios external command definition:
// Desfinition: SCHEDULE_FORCED_SVC_CHECK;<host_name>;<service_description>;<check_time>
// Description:
//  Schedules a forced active check of a particular service at "check_time".  The "check_time" argument is specified in time_t format (seconds since the UNIX epoch).   Forced checks are performed regardless of what time it is (e.g. timeperiod restrictions are ignored) and whether or not active checks are enabled on a service-specific or program-wide basis.
func ScheduleForcedSvcCheck(host_name string, service_description string, check_time time.Time) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SCHEDULE_FORCED_SVC_CHECK")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		CheckTime(c, check_time)
	}
}

// ScheduleHostgroupHostDowntime is generated from the nagios external command definition:
// Desfinition: SCHEDULE_HOSTGROUP_HOST_DOWNTIME;<hostgroup_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
// Description:
//  Schedules downtime for all hosts in a specified hostgroup.  If the "fixed" argument is set to one (1), downtime will start and end at the times specified by the "start" and "end" arguments.  Otherwise, downtime will begin between the "start" and "end" times and last for "duration" seconds.  The "start" and "end" arguments are specified in time_t format (seconds since the UNIX epoch).  The host downtime entries can be triggered by another downtime entry if the "trigger_id" is set to the ID of another scheduled downtime entry.  Set the "trigger_id" argument to zero (0) if the downtime for the hosts should not be triggered by another downtime entry.
func ScheduleHostgroupHostDowntime(hostgroup_name string, start_time time.Time, end_time time.Time, fixed bool, trigger_id int, duration time.Duration, author string, comment string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SCHEDULE_HOSTGROUP_HOST_DOWNTIME")
		HostGroupName(c, hostgroup_name)
		Start(c, start_time)
		End(c, end_time)
		Fixed(c, fixed)
		TriggerID(c, trigger_id)
		Duration(c, duration)
		Author(c, author)
		Comment(c, comment)
	}
}

// ScheduleHostgroupSvcDowntime is generated from the nagios external command definition:
// Desfinition: SCHEDULE_HOSTGROUP_SVC_DOWNTIME;<hostgroup_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
// Description:
//  Schedules downtime for all services associated with hosts in a specified servicegroup.  If the "fixed" argument is set to one (1), downtime will start and end at the times specified by the "start" and "end" arguments.  Otherwise, downtime will begin between the "start" and "end" times and last for "duration" seconds.  The "start" and "end" arguments are specified in time_t format (seconds since the UNIX epoch).  The service downtime entries can be triggered by another downtime entry if the "trigger_id" is set to the ID of another scheduled downtime entry.  Set the "trigger_id" argument to zero (0) if the downtime for the services should not be triggered by another downtime entry.
func ScheduleHostgroupSvcDowntime(hostgroup_name string, start_time time.Time, end_time time.Time, fixed bool, trigger_id int, duration time.Duration, author string, comment string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SCHEDULE_HOSTGROUP_SVC_DOWNTIME")
		HostGroupName(c, hostgroup_name)
		Start(c, start_time)
		End(c, end_time)
		Fixed(c, fixed)
		TriggerID(c, trigger_id)
		Duration(c, duration)
		Author(c, author)
		Comment(c, comment)
	}
}

// ScheduleHostCheck is generated from the nagios external command definition:
// Desfinition: SCHEDULE_HOST_CHECK;<host_name>;<check_time>
// Description:
//  Schedules the next active check of a particular host at "check_time".  The "check_time" argument is specified in time_t format (seconds since the UNIX epoch).  Note that the host may not actually be checked at the time you specify.  This could occur for a number of reasons: active checks are disabled on a program-wide or service-specific basis, the host is already scheduled to be checked at an earlier time, etc.  If you want to force the host check to occur at the time you specify, look at the SCHEDULE_FORCED_HOST_CHECK command.
func ScheduleHostCheck(host_name string, check_time time.Time) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SCHEDULE_HOST_CHECK")
		Hostname(c, host_name)
		CheckTime(c, check_time)
	}
}

// ScheduleHostDowntime is generated from the nagios external command definition:
// Desfinition: SCHEDULE_HOST_DOWNTIME;<host_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
// Description:
//  Schedules downtime for a specified host.  If the "fixed" argument is set to one (1), downtime will start and end at the times specified by the "start" and "end" arguments.  Otherwise, downtime will begin between the "start" and "end" times and last for "duration" seconds.  The "start" and "end" arguments are specified in time_t format (seconds since the UNIX epoch).  The specified host downtime can be triggered by another downtime entry if the "trigger_id" is set to the ID of another scheduled downtime entry.  Set the "trigger_id" argument to zero (0) if the downtime for the specified host should not be triggered by another downtime entry.
func ScheduleHostDowntime(host_name string, start_time time.Time, end_time time.Time, fixed bool, trigger_id int, duration time.Duration, author string, comment string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SCHEDULE_HOST_DOWNTIME")
		Hostname(c, host_name)
		Start(c, start_time)
		End(c, end_time)
		Fixed(c, fixed)
		TriggerID(c, trigger_id)
		Duration(c, duration)
		Author(c, author)
		Comment(c, comment)
	}
}

// ScheduleHostSvcChecks is generated from the nagios external command definition:
// Desfinition: SCHEDULE_HOST_SVC_CHECKS;<host_name>;<check_time>
// Description:
//  Schedules the next active check of all services on a particular host at "check_time".  The "check_time" argument is specified in time_t format (seconds since the UNIX epoch).  Note that the services may not actually be checked at the time you specify.  This could occur for a number of reasons: active checks are disabled on a program-wide or service-specific basis, the services are already scheduled to be checked at an earlier time, etc.  If you want to force the service checks to occur at the time you specify, look at the SCHEDULE_FORCED_HOST_SVC_CHECKS command.
func ScheduleHostSvcChecks(host_name string, check_time time.Time) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SCHEDULE_HOST_SVC_CHECKS")
		Hostname(c, host_name)
		CheckTime(c, check_time)
	}
}

// ScheduleHostSvcDowntime is generated from the nagios external command definition:
// Desfinition: SCHEDULE_HOST_SVC_DOWNTIME;<host_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
// Description:
//  Schedules downtime for all services associated with a particular host.  If the "fixed" argument is set to one (1), downtime will start and end at the times specified by the "start" and "end" arguments.  Otherwise, downtime will begin between the "start" and "end" times and last for "duration" seconds.  The "start" and "end" arguments are specified in time_t format (seconds since the UNIX epoch).  The service downtime entries can be triggered by another downtime entry if the "trigger_id" is set to the ID of another scheduled downtime entry.  Set the "trigger_id" argument to zero (0) if the downtime for the services should not be triggered by another downtime entry.
func ScheduleHostSvcDowntime(host_name string, start_time time.Time, end_time time.Time, fixed bool, trigger_id int, duration time.Duration, author string, comment string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SCHEDULE_HOST_SVC_DOWNTIME")
		Hostname(c, host_name)
		Start(c, start_time)
		End(c, end_time)
		Fixed(c, fixed)
		TriggerID(c, trigger_id)
		Duration(c, duration)
		Author(c, author)
		Comment(c, comment)
	}
}

// ScheduleServicegroupHostDowntime is generated from the nagios external command definition:
// Desfinition: SCHEDULE_SERVICEGROUP_HOST_DOWNTIME;<servicegroup_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
// Description:
//  Schedules downtime for all hosts that have services in a specified servicegroup.  If the "fixed" argument is set to one (1), downtime will start and end at the times specified by the "start" and "end" arguments.  Otherwise, downtime will begin between the "start" and "end" times and last for "duration" seconds.  The "start" and "end" arguments are specified in time_t format (seconds since the UNIX epoch).  The host downtime entries can be triggered by another downtime entry if the "trigger_id" is set to the ID of another scheduled downtime entry.  Set the "trigger_id" argument to zero (0) if the downtime for the hosts should not be triggered by another downtime entry.
func ScheduleServicegroupHostDowntime(servicegroup_name string, start_time time.Time, end_time time.Time, fixed bool, trigger_id int, duration time.Duration, author string, comment string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SCHEDULE_SERVICEGROUP_HOST_DOWNTIME")
		ServiceGroupName(c, servicegroup_name)
		Start(c, start_time)
		End(c, end_time)
		Fixed(c, fixed)
		TriggerID(c, trigger_id)
		Duration(c, duration)
		Author(c, author)
		Comment(c, comment)
	}
}

// ScheduleServicegroupSvcDowntime is generated from the nagios external command definition:
// Desfinition: SCHEDULE_SERVICEGROUP_SVC_DOWNTIME;<servicegroup_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
// Description:
//  Schedules downtime for all services in a specified servicegroup.  If the "fixed" argument is set to one (1), downtime will start and end at the times specified by the "start" and "end" arguments.  Otherwise, downtime will begin between the "start" and "end" times and last for "duration" seconds.  The "start" and "end" arguments are specified in time_t format (seconds since the UNIX epoch).  The service downtime entries can be triggered by another downtime entry if the "trigger_id" is set to the ID of another scheduled downtime entry.  Set the "trigger_id" argument to zero (0) if the downtime for the services should not be triggered by another downtime entry.
func ScheduleServicegroupSvcDowntime(servicegroup_name string, start_time time.Time, end_time time.Time, fixed bool, trigger_id int, duration time.Duration, author string, comment string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SCHEDULE_SERVICEGROUP_SVC_DOWNTIME")
		ServiceGroupName(c, servicegroup_name)
		Start(c, start_time)
		End(c, end_time)
		Fixed(c, fixed)
		TriggerID(c, trigger_id)
		Duration(c, duration)
		Author(c, author)
		Comment(c, comment)
	}
}

// ScheduleSvcCheck is generated from the nagios external command definition:
// Desfinition: SCHEDULE_SVC_CHECK;<host_name>;<service_description>;<check_time>
// Description:
//  Schedules the next active check of a specified service at "check_time".  The "check_time" argument is specified in time_t format (seconds since the UNIX epoch).  Note that the service may not actually be checked at the time you specify.  This could occur for a number of reasons: active checks are disabled on a program-wide or service-specific basis, the service is already scheduled to be checked at an earlier time, etc.  If you want to force the service check to occur at the time you specify, look at the SCHEDULE_FORCED_SVC_CHECK command.
func ScheduleSvcCheck(host_name string, service_description string, check_time time.Time) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SCHEDULE_SVC_CHECK")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		CheckTime(c, check_time)
	}
}

// SendCustomHostNotification is generated from the nagios external command definition:
// Desfinition: SEND_CUSTOM_HOST_NOTIFICATION;<host_name>;<options>;<author>;<comment>
// Description:
//  Allows you to send a custom host notification.  Very useful in dire situations, emergencies or to communicate with all admins that are responsible for a particular host.  When the host notification is sent out, the $NOTIFICATIONTYPE$ macro will be set to "CUSTOM".  The <options> field is a logical OR of the following integer values that affect aspects of the notification that are sent out: 0 = No option (default), 1 = Broadcast (send notification to all normal and all escalated contacts for the host), 2 = Forced (notification is sent out regardless of current time, whether or not notifications are enabled, etc.), 4 = Increment current notification # for the host (this is not done by default for custom notifications).  The comment field can be used with the $NOTIFICATIONCOMMENT$ macro in notification commands.
func SendCustomHostNotification(host_name string, options int, author string, comment string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SEND_CUSTOM_HOST_NOTIFICATION")
		Hostname(c, host_name)
		Options(c, options)
		Author(c, author)
		Comment(c, comment)
	}
}

// SendCustomSvcNotification is generated from the nagios external command definition:
// Desfinition: SEND_CUSTOM_SVC_NOTIFICATION;<host_name>;<service_description>;<options>;<author>;<comment>
// Description:
//  Allows you to send a custom service notification.  Very useful in dire situations, emergencies or to communicate with all admins that are responsible for a particular service.  When the service notification is sent out, the $NOTIFICATIONTYPE$ macro will be set to "CUSTOM".  The <options> field is a logical OR of the following integer values that affect aspects of the notification that are sent out: 0 = No option (default), 1 = Broadcast (send notification to all normal and all escalated contacts for the service), 2 = Forced (notification is sent out regardless of current time, whether or not notifications are enabled, etc.), 4 = Increment current notification # for the service(this is not done by default for custom notifications).   The comment field can be used with the $NOTIFICATIONCOMMENT$ macro in notification commands.
func SendCustomSvcNotification(host_name string, service_description string, options int, author string, comment string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SEND_CUSTOM_SVC_NOTIFICATION")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		Options(c, options)
		Author(c, author)
		Comment(c, comment)
	}
}

// SetHostNotificationNumber is generated from the nagios external command definition:
// Desfinition: SET_HOST_NOTIFICATION_NUMBER;<host_name>;<notification_number>
// Description:
//  Sets the current notification number for a particular host.  A value of 0 indicates that no notification has yet been sent for the current host problem.  Useful for forcing an escalation (based on notification number) or replicating notification information in redundant monitoring environments. Notification numbers greater than zero have no noticeable affect on the notification process if the host is currently in an UP state.
func SetHostNotificationNumber(host_name string, notification_number int) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SET_HOST_NOTIFICATION_NUMBER")
		Hostname(c, host_name)
		NotificationNumber(c, notification_number)
	}
}

// SetSvcNotificationNumber is generated from the nagios external command definition:
// Desfinition: SET_SVC_NOTIFICATION_NUMBER;<host_name>;<service_description>;<notification_number>
// Description:
//  Sets the current notification number for a particular service.  A value of 0 indicates that no notification has yet been sent for the current service problem.  Useful for forcing an escalation (based on notification number) or replicating notification information in redundant monitoring environments. Notification numbers greater than zero have no noticeable affect on the notification process if the service is currently in an OK state.
func SetSvcNotificationNumber(host_name string, service_description string, notification_number int) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SET_SVC_NOTIFICATION_NUMBER")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
		NotificationNumber(c, notification_number)
	}
}

// ShutdownProgram is generated from the nagios external command definition:
// Desfinition: SHUTDOWN_PROGRAM
// Description:
//  Shuts down the Nagios process.
func ShutdownProgram() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("SHUTDOWN_PROGRAM")
	}
}

// StartAcceptingPassiveHostChecks is generated from the nagios external command definition:
// Desfinition: START_ACCEPTING_PASSIVE_HOST_CHECKS
// Description:
//  Enables acceptance and processing of passive host checks on a program-wide basis.
func StartAcceptingPassiveHostChecks() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("START_ACCEPTING_PASSIVE_HOST_CHECKS")
	}
}

// StartAcceptingPassiveSvcChecks is generated from the nagios external command definition:
// Desfinition: START_ACCEPTING_PASSIVE_SVC_CHECKS
// Description:
//  Enables passive service checks on a program-wide basis.
func StartAcceptingPassiveSvcChecks() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("START_ACCEPTING_PASSIVE_SVC_CHECKS")
	}
}

// StartExecutingHostChecks is generated from the nagios external command definition:
// Desfinition: START_EXECUTING_HOST_CHECKS
// Description:
//  Enables active host checks on a program-wide basis.
func StartExecutingHostChecks() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("START_EXECUTING_HOST_CHECKS")
	}
}

// StartExecutingSvcChecks is generated from the nagios external command definition:
// Desfinition: START_EXECUTING_SVC_CHECKS
// Description:
//  Enables active checks of services on a program-wide basis.
func StartExecutingSvcChecks() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("START_EXECUTING_SVC_CHECKS")
	}
}

// StartObsessingOverHost is generated from the nagios external command definition:
// Desfinition: START_OBSESSING_OVER_HOST;<host_name>
// Description:
//  Enables processing of host checks via the OCHP command for the specified host.
func StartObsessingOverHost(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("START_OBSESSING_OVER_HOST")
		Hostname(c, host_name)
	}
}

// StartObsessingOverHostChecks is generated from the nagios external command definition:
// Desfinition: START_OBSESSING_OVER_HOST_CHECKS
// Description:
//  Enables processing of host checks via the OCHP command on a program-wide basis.
func StartObsessingOverHostChecks() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("START_OBSESSING_OVER_HOST_CHECKS")
	}
}

// StartObsessingOverSvc is generated from the nagios external command definition:
// Desfinition: START_OBSESSING_OVER_SVC;<host_name>;<service_description>
// Description:
//  Enables processing of service checks via the OCSP command for the specified service.
func StartObsessingOverSvc(host_name string, service_description string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("START_OBSESSING_OVER_SVC")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
	}
}

// StartObsessingOverSvcChecks is generated from the nagios external command definition:
// Desfinition: START_OBSESSING_OVER_SVC_CHECKS
// Description:
//  Enables processing of service checks via the OCSP command on a program-wide basis.
func StartObsessingOverSvcChecks() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("START_OBSESSING_OVER_SVC_CHECKS")
	}
}

// StopAcceptingPassiveHostChecks is generated from the nagios external command definition:
// Desfinition: STOP_ACCEPTING_PASSIVE_HOST_CHECKS
// Description:
//  Disables acceptance and processing of passive host checks on a program-wide basis.
func StopAcceptingPassiveHostChecks() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("STOP_ACCEPTING_PASSIVE_HOST_CHECKS")
	}
}

// StopAcceptingPassiveSvcChecks is generated from the nagios external command definition:
// Desfinition: STOP_ACCEPTING_PASSIVE_SVC_CHECKS
// Description:
//  Disables passive service checks on a program-wide basis.
func StopAcceptingPassiveSvcChecks() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("STOP_ACCEPTING_PASSIVE_SVC_CHECKS")
	}
}

// StopExecutingHostChecks is generated from the nagios external command definition:
// Desfinition: STOP_EXECUTING_HOST_CHECKS
// Description:
//  Disables active host checks on a program-wide basis.
func StopExecutingHostChecks() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("STOP_EXECUTING_HOST_CHECKS")
	}
}

// StopExecutingSvcChecks is generated from the nagios external command definition:
// Desfinition: STOP_EXECUTING_SVC_CHECKS
// Description:
//  Disables active checks of services on a program-wide basis.
func StopExecutingSvcChecks() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("STOP_EXECUTING_SVC_CHECKS")
	}
}

// StopObsessingOverHost is generated from the nagios external command definition:
// Desfinition: STOP_OBSESSING_OVER_HOST;<host_name>
// Description:
//  Disables processing of host checks via the OCHP command for the specified host.
func StopObsessingOverHost(host_name string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("STOP_OBSESSING_OVER_HOST")
		Hostname(c, host_name)
	}
}

// StopObsessingOverHostChecks is generated from the nagios external command definition:
// Desfinition: STOP_OBSESSING_OVER_HOST_CHECKS
// Description:
//  Disables processing of host checks via the OCHP command on a program-wide basis.
func StopObsessingOverHostChecks() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("STOP_OBSESSING_OVER_HOST_CHECKS")
	}
}

// StopObsessingOverSvc is generated from the nagios external command definition:
// Desfinition: STOP_OBSESSING_OVER_SVC;<host_name>;<service_description>
// Description:
//  Disables processing of service checks via the OCSP command for the specified service.
func StopObsessingOverSvc(host_name string, service_description string) lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("STOP_OBSESSING_OVER_SVC")
		Hostname(c, host_name)
		ServiceDescription(c, service_description)
	}
}

// StopObsessingOverSvcChecks is generated from the nagios external command definition:
// Desfinition: STOP_OBSESSING_OVER_SVC_CHECKS
// Description:
//  Disables processing of service checks via the OCSP command on a program-wide basis.
func StopObsessingOverSvcChecks() lvst.CommandOpFunc {
	return func(c *lvst.Command) {
		c.Raw("STOP_OBSESSING_OVER_SVC_CHECKS")
	}
}
