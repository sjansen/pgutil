package ddl

import (
	"bytes"
)

%%{
	machine trigger;
	write data;
}%%

func ParseTrigger(data string) (*Trigger, error) {
	trigger := &Trigger{}


	buffer := &bytes.Buffer{}
	cs, p, pe, eof := 0, 0, len(data), -1
	%%{
	action buffer_fc       {
		buffer.WriteByte(fc)
	}
	action match_constraint   { trigger.Constraint = true }
	action match_name         {
		trigger.Name = buffer.String()
		buffer.Reset()
	}

	ws = space+;
	ident = [a-zA-Z][_a-zA-Z0-9]*;

	main := space*
		'CREATE'i ws ( 'CONSTRAINT'i@match_constraint ws )?  'TRIGGER'i
		ws ( ident $ buffer_fc % match_name )
		space*
		'\n'?
		;

	write init;
	write exec;
	}%%

	if cs < trigger_first_final {
		return nil, &parseError{
		    cs: cs,
		    data: data,
		}
	}

	return trigger, nil
}
