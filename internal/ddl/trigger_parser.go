package ddl

import (
	"bytes"
)

var trigger_start int = 1
var _ = trigger_start
var trigger_first_final int = 28
var _ = trigger_first_final
var trigger_error int = 0
var _ = trigger_error
var trigger_en_main int = 1
var _ = trigger_en_main

func ParseTrigger(data string) (*Trigger, error) {
	trigger := &Trigger{}

	buffer := &bytes.Buffer{}
	cs, p, pe, eof := 0, 0, len(data), -1
	{
		cs = int(trigger_start)
	}
	{
		if p == pe {
			goto _test_eof

		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29

		}
		goto st_out
	_st1:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof1

		}
	st_case_1:
		switch data[p] {
		case 32:
			{
				goto _st1
			}
		case 67:
			{
				goto _st2
			}
		case 99:
			{
				goto _st2
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st1
			}

		}
		{
			goto _st0
		}
	st_case_0:
	_st0:
		cs = 0
		goto _pop
	_st2:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof2

		}
	st_case_2:
		switch data[p] {
		case 82:
			{
				goto _st3
			}
		case 114:
			{
				goto _st3
			}

		}
		{
			goto _st0
		}
	_st3:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof3

		}
	st_case_3:
		switch data[p] {
		case 69:
			{
				goto _st4
			}
		case 101:
			{
				goto _st4
			}

		}
		{
			goto _st0
		}
	_st4:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof4

		}
	st_case_4:
		switch data[p] {
		case 65:
			{
				goto _st5
			}
		case 97:
			{
				goto _st5
			}

		}
		{
			goto _st0
		}
	_st5:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof5

		}
	st_case_5:
		switch data[p] {
		case 84:
			{
				goto _st6
			}
		case 116:
			{
				goto _st6
			}

		}
		{
			goto _st0
		}
	_st6:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof6

		}
	st_case_6:
		switch data[p] {
		case 69:
			{
				goto _st7
			}
		case 101:
			{
				goto _st7
			}

		}
		{
			goto _st0
		}
	_st7:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof7

		}
	st_case_7:
		if (data[p]) == 32 {
			{
				goto _st8
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st8
			}

		}
		{
			goto _st0
		}
	_st8:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof8

		}
	st_case_8:
		switch data[p] {
		case 32:
			{
				goto _st8
			}
		case 67:
			{
				goto _st9
			}
		case 84:
			{
				goto _st20
			}
		case 99:
			{
				goto _st9
			}
		case 116:
			{
				goto _st20
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st8
			}

		}
		{
			goto _st0
		}
	_st9:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof9

		}
	st_case_9:
		switch data[p] {
		case 79:
			{
				goto _st10
			}
		case 111:
			{
				goto _st10
			}

		}
		{
			goto _st0
		}
	_st10:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof10

		}
	st_case_10:
		switch data[p] {
		case 78:
			{
				goto _st11
			}
		case 110:
			{
				goto _st11
			}

		}
		{
			goto _st0
		}
	_st11:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof11

		}
	st_case_11:
		switch data[p] {
		case 83:
			{
				goto _st12
			}
		case 115:
			{
				goto _st12
			}

		}
		{
			goto _st0
		}
	_st12:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof12

		}
	st_case_12:
		switch data[p] {
		case 84:
			{
				goto _st13
			}
		case 116:
			{
				goto _st13
			}

		}
		{
			goto _st0
		}
	_st13:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof13

		}
	st_case_13:
		switch data[p] {
		case 82:
			{
				goto _st14
			}
		case 114:
			{
				goto _st14
			}

		}
		{
			goto _st0
		}
	_st14:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof14

		}
	st_case_14:
		switch data[p] {
		case 65:
			{
				goto _st15
			}
		case 97:
			{
				goto _st15
			}

		}
		{
			goto _st0
		}
	_st15:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof15

		}
	st_case_15:
		switch data[p] {
		case 73:
			{
				goto _st16
			}
		case 105:
			{
				goto _st16
			}

		}
		{
			goto _st0
		}
	_st16:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof16

		}
	st_case_16:
		switch data[p] {
		case 78:
			{
				goto _st17
			}
		case 110:
			{
				goto _st17
			}

		}
		{
			goto _st0
		}
	_st17:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof17

		}
	st_case_17:
		switch data[p] {
		case 84:
			{
				goto _ctr19
			}
		case 116:
			{
				goto _ctr19
			}

		}
		{
			goto _st0
		}
	_ctr19:
		{
			trigger.Constraint = true
		}

		goto _st18
	_st18:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof18

		}
	st_case_18:
		if (data[p]) == 32 {
			{
				goto _st19
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st19
			}

		}
		{
			goto _st0
		}
	_st19:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof19

		}
	st_case_19:
		switch data[p] {
		case 32:
			{
				goto _st19
			}
		case 84:
			{
				goto _st20
			}
		case 116:
			{
				goto _st20
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st19
			}

		}
		{
			goto _st0
		}
	_st20:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof20

		}
	st_case_20:
		switch data[p] {
		case 82:
			{
				goto _st21
			}
		case 114:
			{
				goto _st21
			}

		}
		{
			goto _st0
		}
	_st21:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof21

		}
	st_case_21:
		switch data[p] {
		case 73:
			{
				goto _st22
			}
		case 105:
			{
				goto _st22
			}

		}
		{
			goto _st0
		}
	_st22:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof22

		}
	st_case_22:
		switch data[p] {
		case 71:
			{
				goto _st23
			}
		case 103:
			{
				goto _st23
			}

		}
		{
			goto _st0
		}
	_st23:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof23

		}
	st_case_23:
		switch data[p] {
		case 71:
			{
				goto _st24
			}
		case 103:
			{
				goto _st24
			}

		}
		{
			goto _st0
		}
	_st24:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof24

		}
	st_case_24:
		switch data[p] {
		case 69:
			{
				goto _st25
			}
		case 101:
			{
				goto _st25
			}

		}
		{
			goto _st0
		}
	_st25:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof25

		}
	st_case_25:
		switch data[p] {
		case 82:
			{
				goto _st26
			}
		case 114:
			{
				goto _st26
			}

		}
		{
			goto _st0
		}
	_st26:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof26

		}
	st_case_26:
		if (data[p]) == 32 {
			{
				goto _st27
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st27
			}

		}
		{
			goto _st0
		}
	_st27:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof27

		}
	st_case_27:
		if (data[p]) == 32 {
			{
				goto _st27
			}

		}
		if (data[p]) < 65 {
			{
				if 9 <= (data[p]) && (data[p]) <= 13 {
					{
						goto _st27
					}

				}
			}

		} else if (data[p]) > 90 {
			{
				if 97 <= (data[p]) && (data[p]) <= 122 {
					{
						goto _ctr29
					}

				}
			}

		} else {
			{
				goto _ctr29
			}

		}
		{
			goto _st0
		}
	_ctr29:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st28
	_ctr30:
		{
			trigger.Name = buffer.String()
			buffer.Reset()
		}

		goto _st28
	_st28:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof28

		}
	st_case_28:
		switch data[p] {
		case 32:
			{
				goto _ctr31
			}
		case 95:
			{
				goto _ctr29
			}

		}
		if (data[p]) < 48 {
			{
				if 9 <= (data[p]) && (data[p]) <= 13 {
					{
						goto _ctr31
					}

				}
			}

		} else if (data[p]) > 57 {
			{
				if (data[p]) > 90 {
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto _ctr29
							}

						}
					}

				} else if (data[p]) >= 65 {
					{
						goto _ctr29
					}

				}
			}

		} else {
			{
				goto _ctr29
			}

		}
		{
			goto _st0
		}
	_ctr31:
		{
			trigger.Name = buffer.String()
			buffer.Reset()
		}

		goto _st29
	_st29:
		if p == eof {
			{
				if cs >= 28 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof29

		}
	st_case_29:
		if (data[p]) == 32 {
			{
				goto _st29
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st29
			}

		}
		{
			goto _st0
		}
	st_out:
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			{
				switch cs {
				case 1:
					{
						break
					}
				case 0:
					{
						break
					}
				case 2:
					{
						break
					}
				case 3:
					{
						break
					}
				case 4:
					{
						break
					}
				case 5:
					{
						break
					}
				case 6:
					{
						break
					}
				case 7:
					{
						break
					}
				case 8:
					{
						break
					}
				case 9:
					{
						break
					}
				case 10:
					{
						break
					}
				case 11:
					{
						break
					}
				case 12:
					{
						break
					}
				case 13:
					{
						break
					}
				case 14:
					{
						break
					}
				case 15:
					{
						break
					}
				case 16:
					{
						break
					}
				case 17:
					{
						break
					}
				case 18:
					{
						break
					}
				case 19:
					{
						break
					}
				case 20:
					{
						break
					}
				case 21:
					{
						break
					}
				case 22:
					{
						break
					}
				case 23:
					{
						break
					}
				case 24:
					{
						break
					}
				case 25:
					{
						break
					}
				case 26:
					{
						break
					}
				case 27:
					{
						break
					}
				case 28:
					{
						break
					}
				case 29:
					{
						break
					}

				}
				switch cs {

				}
				switch cs {
				case 1:
					goto _st1
				case 0:
					goto _st0
				case 2:
					goto _st2
				case 3:
					goto _st3
				case 4:
					goto _st4
				case 5:
					goto _st5
				case 6:
					goto _st6
				case 7:
					goto _st7
				case 8:
					goto _st8
				case 9:
					goto _st9
				case 10:
					goto _st10
				case 11:
					goto _st11
				case 12:
					goto _st12
				case 13:
					goto _st13
				case 14:
					goto _st14
				case 15:
					goto _st15
				case 16:
					goto _st16
				case 17:
					goto _st17
				case 18:
					goto _st18
				case 19:
					goto _st19
				case 20:
					goto _st20
				case 21:
					goto _st21
				case 22:
					goto _st22
				case 23:
					goto _st23
				case 24:
					goto _st24
				case 25:
					goto _st25
				case 26:
					goto _st26
				case 27:
					goto _st27
				case 28:
					goto _ctr30
				case 29:
					goto _st29
				}
			}

		}
		if cs >= 28 {
			goto _out
		}
	_pop:
		{
		}
	_out:
		{
		}
	}
	if cs < trigger_first_final {
		return nil, &parseError{
			cs:   cs,
			data: data,
		}
	}

	return trigger, nil
}
