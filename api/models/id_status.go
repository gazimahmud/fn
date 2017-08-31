package models

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/validate"
)

// TODO get rid of this. id and status are not more coupled than anything else?
// burn it at the stake

/*IDStatus Id status

swagger:model IdStatus
*/
type IDStatus struct {

	/* Unique identifier representing a specific task.

	Read Only: true
	*/
	ID string `json:"id,omitempty"`

	/* States and valid transitions.

	                 +---------+
	       +---------> delayed <----------------+
	                 +----+----+                |
	                      |                     |
	                      |                     |
	                 +----v----+                |
	       +---------> queued  <----------------+
	                 +----+----+                *
	                      |                     *
	                      |               retry * creates new task
	                 +----v----+                *
	                 | running |                *
	                 +--+-+-+--+                |
	          +---------|-|-|-----+-------------+
	      +---|---------+ | +-----|---------+   |
	      |   |           |       |         |   |
	+-----v---^-+      +--v-------^+     +--v---^-+
	| success   |      | cancelled |     |  error |
	+-----------+      +-----------+     +--------+

	* delayed - has a delay.
	* queued - Ready to be consumed when it's turn comes.
	* running - Currently consumed by a runner which will attempt to process it.
	* success - (or complete? success/error is common javascript terminology)
	* error - Something went wrong. In this case more information can be obtained
	  by inspecting the "reason" field.
	  - timeout
	  - killed - forcibly killed by worker due to resource restrictions or access
	    violations.
	  - bad_exit - exited with non-zero status due to program termination/crash.
	* cancelled - cancelled via API. More information in the reason field.
	  - client_request - Request was cancelled by a client.


	Read Only: true
	*/
	Status string `json:"status,omitempty"`
}

// Validate validates this Id status
func (m *IDStatus) Validate(formats strfmt.Registry) error { return m.validateStatus(formats) }

var idStatusTypeStatusPropEnum []interface{}

// prop value enum
func (m *IDStatus) validateStatusEnum(path, location string, value string) error {
	if idStatusTypeStatusPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["delayed","queued","running","success","error","cancelled"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			idStatusTypeStatusPropEnum = append(idStatusTypeStatusPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, idStatusTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IDStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}