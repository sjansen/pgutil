package ddl

// Code generated by ragel; DO NOT EDIT.

var check_start int = 1
var _ = check_start
var check_first_final int = 56
var _ = check_first_final
var check_error int = 0
var _ = check_error
var check_en_main int = 1
var _ = check_en_main

func ParseCheck(data string) (*Check, error) {
	check := &Check{}

	var mark1, mark2 int

	cs, eof := 0, len(data)
	p, pe := 0, eof
	{
		cs = int(check_start)
	}
	{
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
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
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
		case 58:
			goto st_case_58
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
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 59:
			goto st_case_59
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		}
	_st1:
		p += 1
	st_case_1:
		if p == pe {
			goto _out1
		}
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
			goto _st1
		}
		goto _st0
	_st0:
	st_case_0:
		goto _out0
	_st2:
		p += 1
	st_case_2:
		if p == pe {
			goto _out2
		}
		switch data[p] {
		case 72:
			{
				goto _st3
			}
		case 104:
			{
				goto _st3
			}
		}
		goto _st0
	_st3:
		p += 1
	st_case_3:
		if p == pe {
			goto _out3
		}
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
		goto _st0
	_st4:
		p += 1
	st_case_4:
		if p == pe {
			goto _out4
		}
		switch data[p] {
		case 67:
			{
				goto _st5
			}
		case 99:
			{
				goto _st5
			}
		}
		goto _st0
	_st5:
		p += 1
	st_case_5:
		if p == pe {
			goto _out5
		}
		switch data[p] {
		case 75:
			{
				goto _st6
			}
		case 107:
			{
				goto _st6
			}
		}
		goto _st0
	_st6:
		p += 1
	st_case_6:
		if p == pe {
			goto _out6
		}
		switch data[p] {
		case 32:
			{
				goto _st6
			}
		case 40:
			{
				goto _st7
			}
		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			goto _st6
		}
		goto _st0
	_st7:
		p += 1
	st_case_7:
		if p == pe {
			goto _out7
		}
		if (data[p]) == 41 {
			goto _ctr9
		}
		goto _ctr8
	_ctr8:
		{
			mark1 = p
		}
		goto _st8
	_st8:
		p += 1
	st_case_8:
		if p == pe {
			goto _out8
		}
		if (data[p]) == 41 {
			goto _ctr11
		}
		goto _st8
	_ctr9:
		{
			mark1 = p
		}
		{
			mark2 = p
		}
		goto _st56
	_ctr11:
		{
			mark2 = p
		}
		goto _st56
	_st56:
		p += 1
	st_case_56:
		if p == pe {
			goto _out56
		}
		switch data[p] {
		case 32:
			{
				goto _st57
			}
		case 41:
			{
				goto _ctr11
			}
		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			goto _st57
		}
		goto _st8
	_st57:
		p += 1
	st_case_57:
		if p == pe {
			goto _out57
		}
		switch data[p] {
		case 32:
			{
				goto _st57
			}
		case 41:
			{
				goto _ctr11
			}
		case 68:
			{
				goto _st9
			}
		case 73:
			{
				goto _st19
			}
		case 78:
			{
				goto _st44
			}
		case 100:
			{
				goto _st9
			}
		case 105:
			{
				goto _st19
			}
		case 110:
			{
				goto _st44
			}
		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			goto _st57
		}
		goto _st8
	_st9:
		p += 1
	st_case_9:
		if p == pe {
			goto _out9
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 69:
			{
				goto _st10
			}
		case 101:
			{
				goto _st10
			}
		}
		goto _st8
	_st10:
		p += 1
	st_case_10:
		if p == pe {
			goto _out10
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 70:
			{
				goto _st11
			}
		case 102:
			{
				goto _st11
			}
		}
		goto _st8
	_st11:
		p += 1
	st_case_11:
		if p == pe {
			goto _out11
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 69:
			{
				goto _st12
			}
		case 101:
			{
				goto _st12
			}
		}
		goto _st8
	_st12:
		p += 1
	st_case_12:
		if p == pe {
			goto _out12
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 82:
			{
				goto _st13
			}
		case 114:
			{
				goto _st13
			}
		}
		goto _st8
	_st13:
		p += 1
	st_case_13:
		if p == pe {
			goto _out13
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 82:
			{
				goto _st14
			}
		case 114:
			{
				goto _st14
			}
		}
		goto _st8
	_st14:
		p += 1
	st_case_14:
		if p == pe {
			goto _out14
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 65:
			{
				goto _st15
			}
		case 97:
			{
				goto _st15
			}
		}
		goto _st8
	_st15:
		p += 1
	st_case_15:
		if p == pe {
			goto _out15
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 66:
			{
				goto _st16
			}
		case 98:
			{
				goto _st16
			}
		}
		goto _st8
	_st16:
		p += 1
	st_case_16:
		if p == pe {
			goto _out16
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 76:
			{
				goto _st17
			}
		case 108:
			{
				goto _st17
			}
		}
		goto _st8
	_st17:
		p += 1
	st_case_17:
		if p == pe {
			goto _out17
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 69:
			{
				goto _st18
			}
		case 101:
			{
				goto _st18
			}
		}
		goto _st8
	_st18:
		p += 1
	st_case_18:
		if p == pe {
			goto _out18
		}
		switch data[p] {
		case 32:
			{
				goto _ctr22
			}
		case 41:
			{
				goto _ctr11
			}
		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			goto _ctr22
		}
		goto _st8
	_ctr22:
		{
			check.Deferrable = true
		}
		goto _st58
	_st58:
		p += 1
	st_case_58:
		if p == pe {
			goto _out58
		}
		switch data[p] {
		case 32:
			{
				goto _ctr22
			}
		case 41:
			{
				goto _ctr11
			}
		case 73:
			{
				goto _st19
			}
		case 105:
			{
				goto _st19
			}
		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			goto _ctr22
		}
		goto _st8
	_st19:
		p += 1
	st_case_19:
		if p == pe {
			goto _out19
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 78:
			{
				goto _st20
			}
		case 110:
			{
				goto _st20
			}
		}
		goto _st8
	_st20:
		p += 1
	st_case_20:
		if p == pe {
			goto _out20
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 73:
			{
				goto _st21
			}
		case 105:
			{
				goto _st21
			}
		}
		goto _st8
	_st21:
		p += 1
	st_case_21:
		if p == pe {
			goto _out21
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 84:
			{
				goto _st22
			}
		case 116:
			{
				goto _st22
			}
		}
		goto _st8
	_st22:
		p += 1
	st_case_22:
		if p == pe {
			goto _out22
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 73:
			{
				goto _st23
			}
		case 105:
			{
				goto _st23
			}
		}
		goto _st8
	_st23:
		p += 1
	st_case_23:
		if p == pe {
			goto _out23
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 65:
			{
				goto _st24
			}
		case 97:
			{
				goto _st24
			}
		}
		goto _st8
	_st24:
		p += 1
	st_case_24:
		if p == pe {
			goto _out24
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 76:
			{
				goto _st25
			}
		case 108:
			{
				goto _st25
			}
		}
		goto _st8
	_st25:
		p += 1
	st_case_25:
		if p == pe {
			goto _out25
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 76:
			{
				goto _st26
			}
		case 108:
			{
				goto _st26
			}
		}
		goto _st8
	_st26:
		p += 1
	st_case_26:
		if p == pe {
			goto _out26
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 89:
			{
				goto _st27
			}
		case 121:
			{
				goto _st27
			}
		}
		goto _st8
	_st27:
		p += 1
	st_case_27:
		if p == pe {
			goto _out27
		}
		switch data[p] {
		case 32:
			{
				goto _st28
			}
		case 41:
			{
				goto _ctr11
			}
		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			goto _st28
		}
		goto _st8
	_st28:
		p += 1
	st_case_28:
		if p == pe {
			goto _out28
		}
		switch data[p] {
		case 32:
			{
				goto _st28
			}
		case 41:
			{
				goto _ctr11
			}
		case 68:
			{
				goto _st29
			}
		case 73:
			{
				goto _st36
			}
		case 100:
			{
				goto _st29
			}
		case 105:
			{
				goto _st36
			}
		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			goto _st28
		}
		goto _st8
	_st29:
		p += 1
	st_case_29:
		if p == pe {
			goto _out29
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 69:
			{
				goto _st30
			}
		case 101:
			{
				goto _st30
			}
		}
		goto _st8
	_st30:
		p += 1
	st_case_30:
		if p == pe {
			goto _out30
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 70:
			{
				goto _st31
			}
		case 102:
			{
				goto _st31
			}
		}
		goto _st8
	_st31:
		p += 1
	st_case_31:
		if p == pe {
			goto _out31
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 69:
			{
				goto _st32
			}
		case 101:
			{
				goto _st32
			}
		}
		goto _st8
	_st32:
		p += 1
	st_case_32:
		if p == pe {
			goto _out32
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 82:
			{
				goto _st33
			}
		case 114:
			{
				goto _st33
			}
		}
		goto _st8
	_st33:
		p += 1
	st_case_33:
		if p == pe {
			goto _out33
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 82:
			{
				goto _st34
			}
		case 114:
			{
				goto _st34
			}
		}
		goto _st8
	_st34:
		p += 1
	st_case_34:
		if p == pe {
			goto _out34
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 69:
			{
				goto _st35
			}
		case 101:
			{
				goto _st35
			}
		}
		goto _st8
	_st35:
		p += 1
	st_case_35:
		if p == pe {
			goto _out35
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 68:
			{
				goto _ctr41
			}
		case 100:
			{
				goto _ctr41
			}
		}
		goto _st8
	_ctr41:
		{
			check.InitiallyDeferred = true
		}
		goto _st59
	_st59:
		p += 1
	st_case_59:
		if p == pe {
			goto _out59
		}
		switch data[p] {
		case 32:
			{
				goto _st59
			}
		case 41:
			{
				goto _ctr11
			}
		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			goto _st59
		}
		goto _st8
	_st36:
		p += 1
	st_case_36:
		if p == pe {
			goto _out36
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 77:
			{
				goto _st37
			}
		case 109:
			{
				goto _st37
			}
		}
		goto _st8
	_st37:
		p += 1
	st_case_37:
		if p == pe {
			goto _out37
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 77:
			{
				goto _st38
			}
		case 109:
			{
				goto _st38
			}
		}
		goto _st8
	_st38:
		p += 1
	st_case_38:
		if p == pe {
			goto _out38
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 69:
			{
				goto _st39
			}
		case 101:
			{
				goto _st39
			}
		}
		goto _st8
	_st39:
		p += 1
	st_case_39:
		if p == pe {
			goto _out39
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 68:
			{
				goto _st40
			}
		case 100:
			{
				goto _st40
			}
		}
		goto _st8
	_st40:
		p += 1
	st_case_40:
		if p == pe {
			goto _out40
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 73:
			{
				goto _st41
			}
		case 105:
			{
				goto _st41
			}
		}
		goto _st8
	_st41:
		p += 1
	st_case_41:
		if p == pe {
			goto _out41
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 65:
			{
				goto _st42
			}
		case 97:
			{
				goto _st42
			}
		}
		goto _st8
	_st42:
		p += 1
	st_case_42:
		if p == pe {
			goto _out42
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 84:
			{
				goto _st43
			}
		case 116:
			{
				goto _st43
			}
		}
		goto _st8
	_st43:
		p += 1
	st_case_43:
		if p == pe {
			goto _out43
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 69:
			{
				goto _st59
			}
		case 101:
			{
				goto _st59
			}
		}
		goto _st8
	_st44:
		p += 1
	st_case_44:
		if p == pe {
			goto _out44
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 79:
			{
				goto _st45
			}
		case 111:
			{
				goto _st45
			}
		}
		goto _st8
	_st45:
		p += 1
	st_case_45:
		if p == pe {
			goto _out45
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 84:
			{
				goto _st46
			}
		case 116:
			{
				goto _st46
			}
		}
		goto _st8
	_st46:
		p += 1
	st_case_46:
		if p == pe {
			goto _out46
		}
		switch data[p] {
		case 32:
			{
				goto _st47
			}
		case 41:
			{
				goto _ctr11
			}
		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			goto _st47
		}
		goto _st8
	_st47:
		p += 1
	st_case_47:
		if p == pe {
			goto _out47
		}
		switch data[p] {
		case 32:
			{
				goto _st47
			}
		case 41:
			{
				goto _ctr11
			}
		case 68:
			{
				goto _st48
			}
		case 100:
			{
				goto _st48
			}
		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			goto _st47
		}
		goto _st8
	_st48:
		p += 1
	st_case_48:
		if p == pe {
			goto _out48
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 69:
			{
				goto _st49
			}
		case 101:
			{
				goto _st49
			}
		}
		goto _st8
	_st49:
		p += 1
	st_case_49:
		if p == pe {
			goto _out49
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 70:
			{
				goto _st50
			}
		case 102:
			{
				goto _st50
			}
		}
		goto _st8
	_st50:
		p += 1
	st_case_50:
		if p == pe {
			goto _out50
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 69:
			{
				goto _st51
			}
		case 101:
			{
				goto _st51
			}
		}
		goto _st8
	_st51:
		p += 1
	st_case_51:
		if p == pe {
			goto _out51
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 82:
			{
				goto _st52
			}
		case 114:
			{
				goto _st52
			}
		}
		goto _st8
	_st52:
		p += 1
	st_case_52:
		if p == pe {
			goto _out52
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 82:
			{
				goto _st53
			}
		case 114:
			{
				goto _st53
			}
		}
		goto _st8
	_st53:
		p += 1
	st_case_53:
		if p == pe {
			goto _out53
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 65:
			{
				goto _st54
			}
		case 97:
			{
				goto _st54
			}
		}
		goto _st8
	_st54:
		p += 1
	st_case_54:
		if p == pe {
			goto _out54
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 66:
			{
				goto _st55
			}
		case 98:
			{
				goto _st55
			}
		}
		goto _st8
	_st55:
		p += 1
	st_case_55:
		if p == pe {
			goto _out55
		}
		switch data[p] {
		case 41:
			{
				goto _ctr11
			}
		case 76:
			{
				goto _st43
			}
		case 108:
			{
				goto _st43
			}
		}
		goto _st8
	_out1:
		cs = 1
		goto _out
	_out0:
		cs = 0
		goto _out
	_out2:
		cs = 2
		goto _out
	_out3:
		cs = 3
		goto _out
	_out4:
		cs = 4
		goto _out
	_out5:
		cs = 5
		goto _out
	_out6:
		cs = 6
		goto _out
	_out7:
		cs = 7
		goto _out
	_out8:
		cs = 8
		goto _out
	_out56:
		cs = 56
		goto _out
	_out57:
		cs = 57
		goto _out
	_out9:
		cs = 9
		goto _out
	_out10:
		cs = 10
		goto _out
	_out11:
		cs = 11
		goto _out
	_out12:
		cs = 12
		goto _out
	_out13:
		cs = 13
		goto _out
	_out14:
		cs = 14
		goto _out
	_out15:
		cs = 15
		goto _out
	_out16:
		cs = 16
		goto _out
	_out17:
		cs = 17
		goto _out
	_out18:
		cs = 18
		goto _out
	_out58:
		cs = 58
		goto _out
	_out19:
		cs = 19
		goto _out
	_out20:
		cs = 20
		goto _out
	_out21:
		cs = 21
		goto _out
	_out22:
		cs = 22
		goto _out
	_out23:
		cs = 23
		goto _out
	_out24:
		cs = 24
		goto _out
	_out25:
		cs = 25
		goto _out
	_out26:
		cs = 26
		goto _out
	_out27:
		cs = 27
		goto _out
	_out28:
		cs = 28
		goto _out
	_out29:
		cs = 29
		goto _out
	_out30:
		cs = 30
		goto _out
	_out31:
		cs = 31
		goto _out
	_out32:
		cs = 32
		goto _out
	_out33:
		cs = 33
		goto _out
	_out34:
		cs = 34
		goto _out
	_out35:
		cs = 35
		goto _out
	_out59:
		cs = 59
		goto _out
	_out36:
		cs = 36
		goto _out
	_out37:
		cs = 37
		goto _out
	_out38:
		cs = 38
		goto _out
	_out39:
		cs = 39
		goto _out
	_out40:
		cs = 40
		goto _out
	_out41:
		cs = 41
		goto _out
	_out42:
		cs = 42
		goto _out
	_out43:
		cs = 43
		goto _out
	_out44:
		cs = 44
		goto _out
	_out45:
		cs = 45
		goto _out
	_out46:
		cs = 46
		goto _out
	_out47:
		cs = 47
		goto _out
	_out48:
		cs = 48
		goto _out
	_out49:
		cs = 49
		goto _out
	_out50:
		cs = 50
		goto _out
	_out51:
		cs = 51
		goto _out
	_out52:
		cs = 52
		goto _out
	_out53:
		cs = 53
		goto _out
	_out54:
		cs = 54
		goto _out
	_out55:
		cs = 55
		goto _out
	_out:
		{
		}
	}
	if cs < check_first_final {
		return nil, &parseError{
			cs:   cs,
			data: data,
		}
	}

	check.setExpression(data[mark1:mark2])

	return check, nil
}
