package ddl

// Code generated by go generate; DO NOT EDIT.

import (
	"bytes"
)

%%{
	machine index;
	write data;
}%%

func ParseIndex(data string) (*Index, error) {
	index := &Index{}

	buffer := &bytes.Buffer{}
	cs, p, pe := 0, 0, len(data)
	%%{

	action addToBuffer       {
		buffer.WriteByte(fc)
	}

	action addColumn {
		index.addColumn(buffer.String())
		buffer.Reset()
	}
	action setName {
		index.setName(buffer.String())
		buffer.Reset()
	}
	action setTable {
		index.setTable(buffer.String())
		buffer.Reset()
	}
	action setUsing {
		index.setUsing(buffer.String())
		buffer.Reset()
	}

	action matchUnique   { index.Unique = true }

	ws = space+;
	ident = [a-zA-Z][_a-zA-Z0-9]*;

	main := space*
		'CREATE'i ws ( 'UNIQUE'i @ matchUnique ws )?  'INDEX'i
		ws ( ident $ addToBuffer % setName )
		ws 'ON'i ws ( ident $ addToBuffer % setTable )
		(ws 'USING'i ws ( ident $ addToBuffer % setUsing ))?
		ws '('
		  ( ident $ addToBuffer % addColumn )
		  ( ws? ',' ws? ( ident $ addToBuffer % addColumn ) )*
		')'
		space*
		;

	write init;
	write exec;
	}%%

	if cs < index_first_final {
		return nil, &parseError{
		    cs: cs,
		    data: data,
		}
	}

	return index, nil
}
