// Code generated by goyacc -o generated.go grammar.y. DO NOT EDIT.

//line grammar.y:2
package sqlparser

import __yyfmt__ "fmt"

//line grammar.y:2

import (
	"github.com/sjansen/pgutil/internal/sql"
)

type option struct {
	Name  string
	Value interface{}
}

func newOption(name string, value interface{}) *option {
	return &option{
		Name:  name,
		Value: value,
	}
}

func newOptionList(opt *option) []*option {
	slice := make([]*option, 0, 4)
	slice = append(slice, opt)
	return slice
}

//line grammar.y:33
type yySymType struct {
	yys  int
	ast  interface{}
	bool bool
	opt  *option
	opts []*option
	str  string
}

const ABORT = 57346
const ABSOLUTE = 57347
const ACCESS = 57348
const ACTION = 57349
const ADD = 57350
const ADMIN = 57351
const AFTER = 57352
const AGGREGATE = 57353
const ALL = 57354
const ALSO = 57355
const ALTER = 57356
const ALWAYS = 57357
const ANALYSE = 57358
const ANALYZE = 57359
const AND = 57360
const ANY = 57361
const ARRAY = 57362
const AS = 57363
const ASC = 57364
const ASSERTION = 57365
const ASSIGNMENT = 57366
const ASYMMETRIC = 57367
const AT = 57368
const ATTACH = 57369
const ATTRIBUTE = 57370
const AUTHORIZATION = 57371
const BACKWARD = 57372
const BEFORE = 57373
const BEGIN = 57374
const BETWEEN = 57375
const BIGINT = 57376
const BINARY = 57377
const BIT = 57378
const BOOLEAN = 57379
const BOTH = 57380
const BY = 57381
const CACHE = 57382
const CALL = 57383
const CALLED = 57384
const CASCADE = 57385
const CASCADED = 57386
const CASE = 57387
const CAST = 57388
const CATALOG = 57389
const CHAIN = 57390
const CHAR = 57391
const CHARACTER = 57392
const CHARACTERISTICS = 57393
const CHECK = 57394
const CHECKPOINT = 57395
const CLASS = 57396
const CLOSE = 57397
const CLUSTER = 57398
const COALESCE = 57399
const COLLATE = 57400
const COLLATION = 57401
const COLUMN = 57402
const COLUMNS = 57403
const COMMENT = 57404
const COMMENTS = 57405
const COMMIT = 57406
const COMMITTED = 57407
const CONCURRENTLY = 57408
const CONFIGURATION = 57409
const CONFLICT = 57410
const CONNECTION = 57411
const CONSTRAINT = 57412
const CONSTRAINTS = 57413
const CONTENT = 57414
const CONTINUE = 57415
const CONVERSION = 57416
const COPY = 57417
const COST = 57418
const CREATE = 57419
const CROSS = 57420
const CSV = 57421
const CUBE = 57422
const CURRENT = 57423
const CURRENT_CATALOG = 57424
const CURRENT_DATE = 57425
const CURRENT_ROLE = 57426
const CURRENT_SCHEMA = 57427
const CURRENT_TIME = 57428
const CURRENT_TIMESTAMP = 57429
const CURRENT_USER = 57430
const CURSOR = 57431
const CYCLE = 57432
const DATA = 57433
const DATABASE = 57434
const DAY = 57435
const DEALLOCATE = 57436
const DEC = 57437
const DECIMAL = 57438
const DECLARE = 57439
const DEFAULT = 57440
const DEFAULTS = 57441
const DEFERRABLE = 57442
const DEFERRED = 57443
const DEFINER = 57444
const DELETE = 57445
const DELIMITER = 57446
const DELIMITERS = 57447
const DEPENDS = 57448
const DESC = 57449
const DETACH = 57450
const DICTIONARY = 57451
const DISABLE = 57452
const DISCARD = 57453
const DISTINCT = 57454
const DO = 57455
const DOCUMENT = 57456
const DOMAIN = 57457
const DOUBLE = 57458
const DROP = 57459
const EACH = 57460
const ELSE = 57461
const ENABLE = 57462
const ENCODING = 57463
const ENCRYPTED = 57464
const END = 57465
const ENUM = 57466
const ESCAPE = 57467
const EVENT = 57468
const EXCEPT = 57469
const EXCLUDE = 57470
const EXCLUDING = 57471
const EXCLUSIVE = 57472
const EXECUTE = 57473
const EXISTS = 57474
const EXPLAIN = 57475
const EXPRESSION = 57476
const EXTENSION = 57477
const EXTERNAL = 57478
const EXTRACT = 57479
const FALSE = 57480
const FAMILY = 57481
const FETCH = 57482
const FILTER = 57483
const FIRST = 57484
const FLOAT = 57485
const FOLLOWING = 57486
const FOR = 57487
const FORCE = 57488
const FOREIGN = 57489
const FORWARD = 57490
const FREEZE = 57491
const FROM = 57492
const FULL = 57493
const FUNCTION = 57494
const FUNCTIONS = 57495
const GENERATED = 57496
const GLOBAL = 57497
const GRANT = 57498
const GRANTED = 57499
const GREATEST = 57500
const GROUP = 57501
const GROUPING = 57502
const GROUPS = 57503
const HANDLER = 57504
const HAVING = 57505
const HEADER = 57506
const HOLD = 57507
const HOUR = 57508
const IDENTITY = 57509
const IF = 57510
const ILIKE = 57511
const IMMEDIATE = 57512
const IMMUTABLE = 57513
const IMPLICIT = 57514
const IMPORT = 57515
const IN = 57516
const INCLUDE = 57517
const INCLUDING = 57518
const INCREMENT = 57519
const INDEX = 57520
const INDEXES = 57521
const INHERIT = 57522
const INHERITS = 57523
const INITIALLY = 57524
const INLINE = 57525
const INNER = 57526
const INOUT = 57527
const INPUT = 57528
const INSENSITIVE = 57529
const INSERT = 57530
const INSTEAD = 57531
const INT = 57532
const INTEGER = 57533
const INTERSECT = 57534
const INTERVAL = 57535
const INTO = 57536
const INVOKER = 57537
const IS = 57538
const ISNULL = 57539
const ISOLATION = 57540
const JOIN = 57541
const KEY = 57542
const LABEL = 57543
const LANGUAGE = 57544
const LARGE = 57545
const LAST = 57546
const LATERAL = 57547
const LEADING = 57548
const LEAKPROOF = 57549
const LEAST = 57550
const LEFT = 57551
const LEVEL = 57552
const LIKE = 57553
const LIMIT = 57554
const LISTEN = 57555
const LOAD = 57556
const LOCAL = 57557
const LOCALTIME = 57558
const LOCALTIMESTAMP = 57559
const LOCATION = 57560
const LOCK = 57561
const LOCKED = 57562
const LOGGED = 57563
const MAPPING = 57564
const MATCH = 57565
const MATERIALIZED = 57566
const MAXVALUE = 57567
const METHOD = 57568
const MINUTE = 57569
const MINVALUE = 57570
const MODE = 57571
const MONTH = 57572
const MOVE = 57573
const NAME = 57574
const NAMES = 57575
const NATIONAL = 57576
const NATURAL = 57577
const NCHAR = 57578
const NEW = 57579
const NEXT = 57580
const NFC = 57581
const NFD = 57582
const NFKC = 57583
const NFKD = 57584
const NO = 57585
const NONE = 57586
const NORMALIZE = 57587
const NORMALIZED = 57588
const NOT = 57589
const NOTHING = 57590
const NOTIFY = 57591
const NOTNULL = 57592
const NOWAIT = 57593
const NULL = 57594
const NULLIF = 57595
const NULLS = 57596
const NUMERIC = 57597
const OBJECT = 57598
const OF = 57599
const OFF = 57600
const OFFSET = 57601
const OIDS = 57602
const OLD = 57603
const ON = 57604
const ONLY = 57605
const OPERATOR = 57606
const OPTION = 57607
const OPTIONS = 57608
const OR = 57609
const ORDER = 57610
const ORDINALITY = 57611
const OTHERS = 57612
const OUT = 57613
const OUTER = 57614
const OVER = 57615
const OVERLAPS = 57616
const OVERLAY = 57617
const OVERRIDING = 57618
const OWNED = 57619
const OWNER = 57620
const PARALLEL = 57621
const PARSER = 57622
const PARTIAL = 57623
const PARTITION = 57624
const PASSING = 57625
const PASSWORD = 57626
const PLACING = 57627
const PLANS = 57628
const POLICY = 57629
const POSITION = 57630
const PRECEDING = 57631
const PRECISION = 57632
const PRESERVE = 57633
const PREPARE = 57634
const PREPARED = 57635
const PRIMARY = 57636
const PRIOR = 57637
const PRIVILEGES = 57638
const PROCEDURAL = 57639
const PROCEDURE = 57640
const PROCEDURES = 57641
const PROGRAM = 57642
const PUBLICATION = 57643
const QUOTE = 57644
const RANGE = 57645
const READ = 57646
const REAL = 57647
const REASSIGN = 57648
const RECHECK = 57649
const RECURSIVE = 57650
const REF = 57651
const REFERENCES = 57652
const REFERENCING = 57653
const REFRESH = 57654
const REINDEX = 57655
const RELATIVE = 57656
const RELEASE = 57657
const RENAME = 57658
const REPEATABLE = 57659
const REPLACE = 57660
const REPLICA = 57661
const RESET = 57662
const RESTART = 57663
const RESTRICT = 57664
const RETURNING = 57665
const RETURNS = 57666
const REVOKE = 57667
const RIGHT = 57668
const ROLE = 57669
const ROLLBACK = 57670
const ROLLUP = 57671
const ROUTINE = 57672
const ROUTINES = 57673
const ROW = 57674
const ROWS = 57675
const RULE = 57676
const SAVEPOINT = 57677
const SCHEMA = 57678
const SCHEMAS = 57679
const SCROLL = 57680
const SEARCH = 57681
const SECOND = 57682
const SECURITY = 57683
const SELECT = 57684
const SEQUENCE = 57685
const SEQUENCES = 57686
const SERIALIZABLE = 57687
const SERVER = 57688
const SESSION = 57689
const SESSION_USER = 57690
const SET = 57691
const SETS = 57692
const SETOF = 57693
const SHARE = 57694
const SHOW = 57695
const SIMILAR = 57696
const SIMPLE = 57697
const SKIP = 57698
const SMALLINT = 57699
const SNAPSHOT = 57700
const SOME = 57701
const SQL = 57702
const STABLE = 57703
const STANDALONE = 57704
const START = 57705
const STATEMENT = 57706
const STATISTICS = 57707
const STDIN = 57708
const STDOUT = 57709
const STORAGE = 57710
const STORED = 57711
const STRICT = 57712
const STRIP = 57713
const SUBSCRIPTION = 57714
const SUBSTRING = 57715
const SUPPORT = 57716
const SYMMETRIC = 57717
const SYSID = 57718
const SYSTEM = 57719
const TABLE = 57720
const TABLES = 57721
const TABLESAMPLE = 57722
const TABLESPACE = 57723
const TEMP = 57724
const TEMPLATE = 57725
const TEMPORARY = 57726
const TEXT = 57727
const THEN = 57728
const TIES = 57729
const TIME = 57730
const TIMESTAMP = 57731
const TO = 57732
const TRAILING = 57733
const TRANSACTION = 57734
const TRANSFORM = 57735
const TREAT = 57736
const TRIGGER = 57737
const TRIM = 57738
const TRUE = 57739
const TRUNCATE = 57740
const TRUSTED = 57741
const TYPE = 57742
const TYPES = 57743
const UESCAPE = 57744
const UNBOUNDED = 57745
const UNCOMMITTED = 57746
const UNENCRYPTED = 57747
const UNION = 57748
const UNIQUE = 57749
const UNKNOWN = 57750
const UNLISTEN = 57751
const UNLOGGED = 57752
const UNTIL = 57753
const UPDATE = 57754
const USER = 57755
const USING = 57756
const VACUUM = 57757
const VALID = 57758
const VALIDATE = 57759
const VALIDATOR = 57760
const VALUE = 57761
const VALUES = 57762
const VARCHAR = 57763
const VARIADIC = 57764
const VARYING = 57765
const VERBOSE = 57766
const VERSION = 57767
const VIEW = 57768
const VIEWS = 57769
const VOLATILE = 57770
const WHEN = 57771
const WHERE = 57772
const WHITESPACE = 57773
const WINDOW = 57774
const WITH = 57775
const WITHIN = 57776
const WITHOUT = 57777
const WORK = 57778
const WRAPPER = 57779
const WRITE = 57780
const XML = 57781
const XMLATTRIBUTES = 57782
const XMLCONCAT = 57783
const XMLELEMENT = 57784
const XMLEXISTS = 57785
const XMLFOREST = 57786
const XMLNAMESPACES = 57787
const XMLPARSE = 57788
const XMLPI = 57789
const XMLROOT = 57790
const XMLSERIALIZE = 57791
const XMLTABLE = 57792
const YEAR = 57793
const YES = 57794
const ZONE = 57795
const UNEXPECTED_SYMBOL = 57796
const Identifier = 57797
const Name = 57798

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ABORT",
	"ABSOLUTE",
	"ACCESS",
	"ACTION",
	"ADD",
	"ADMIN",
	"AFTER",
	"AGGREGATE",
	"ALL",
	"ALSO",
	"ALTER",
	"ALWAYS",
	"ANALYSE",
	"ANALYZE",
	"AND",
	"ANY",
	"ARRAY",
	"AS",
	"ASC",
	"ASSERTION",
	"ASSIGNMENT",
	"ASYMMETRIC",
	"AT",
	"ATTACH",
	"ATTRIBUTE",
	"AUTHORIZATION",
	"BACKWARD",
	"BEFORE",
	"BEGIN",
	"BETWEEN",
	"BIGINT",
	"BINARY",
	"BIT",
	"BOOLEAN",
	"BOTH",
	"BY",
	"CACHE",
	"CALL",
	"CALLED",
	"CASCADE",
	"CASCADED",
	"CASE",
	"CAST",
	"CATALOG",
	"CHAIN",
	"CHAR",
	"CHARACTER",
	"CHARACTERISTICS",
	"CHECK",
	"CHECKPOINT",
	"CLASS",
	"CLOSE",
	"CLUSTER",
	"COALESCE",
	"COLLATE",
	"COLLATION",
	"COLUMN",
	"COLUMNS",
	"COMMENT",
	"COMMENTS",
	"COMMIT",
	"COMMITTED",
	"CONCURRENTLY",
	"CONFIGURATION",
	"CONFLICT",
	"CONNECTION",
	"CONSTRAINT",
	"CONSTRAINTS",
	"CONTENT",
	"CONTINUE",
	"CONVERSION",
	"COPY",
	"COST",
	"CREATE",
	"CROSS",
	"CSV",
	"CUBE",
	"CURRENT",
	"CURRENT_CATALOG",
	"CURRENT_DATE",
	"CURRENT_ROLE",
	"CURRENT_SCHEMA",
	"CURRENT_TIME",
	"CURRENT_TIMESTAMP",
	"CURRENT_USER",
	"CURSOR",
	"CYCLE",
	"DATA",
	"DATABASE",
	"DAY",
	"DEALLOCATE",
	"DEC",
	"DECIMAL",
	"DECLARE",
	"DEFAULT",
	"DEFAULTS",
	"DEFERRABLE",
	"DEFERRED",
	"DEFINER",
	"DELETE",
	"DELIMITER",
	"DELIMITERS",
	"DEPENDS",
	"DESC",
	"DETACH",
	"DICTIONARY",
	"DISABLE",
	"DISCARD",
	"DISTINCT",
	"DO",
	"DOCUMENT",
	"DOMAIN",
	"DOUBLE",
	"DROP",
	"EACH",
	"ELSE",
	"ENABLE",
	"ENCODING",
	"ENCRYPTED",
	"END",
	"ENUM",
	"ESCAPE",
	"EVENT",
	"EXCEPT",
	"EXCLUDE",
	"EXCLUDING",
	"EXCLUSIVE",
	"EXECUTE",
	"EXISTS",
	"EXPLAIN",
	"EXPRESSION",
	"EXTENSION",
	"EXTERNAL",
	"EXTRACT",
	"FALSE",
	"FAMILY",
	"FETCH",
	"FILTER",
	"FIRST",
	"FLOAT",
	"FOLLOWING",
	"FOR",
	"FORCE",
	"FOREIGN",
	"FORWARD",
	"FREEZE",
	"FROM",
	"FULL",
	"FUNCTION",
	"FUNCTIONS",
	"GENERATED",
	"GLOBAL",
	"GRANT",
	"GRANTED",
	"GREATEST",
	"GROUP",
	"GROUPING",
	"GROUPS",
	"HANDLER",
	"HAVING",
	"HEADER",
	"HOLD",
	"HOUR",
	"IDENTITY",
	"IF",
	"ILIKE",
	"IMMEDIATE",
	"IMMUTABLE",
	"IMPLICIT",
	"IMPORT",
	"IN",
	"INCLUDE",
	"INCLUDING",
	"INCREMENT",
	"INDEX",
	"INDEXES",
	"INHERIT",
	"INHERITS",
	"INITIALLY",
	"INLINE",
	"INNER",
	"INOUT",
	"INPUT",
	"INSENSITIVE",
	"INSERT",
	"INSTEAD",
	"INT",
	"INTEGER",
	"INTERSECT",
	"INTERVAL",
	"INTO",
	"INVOKER",
	"IS",
	"ISNULL",
	"ISOLATION",
	"JOIN",
	"KEY",
	"LABEL",
	"LANGUAGE",
	"LARGE",
	"LAST",
	"LATERAL",
	"LEADING",
	"LEAKPROOF",
	"LEAST",
	"LEFT",
	"LEVEL",
	"LIKE",
	"LIMIT",
	"LISTEN",
	"LOAD",
	"LOCAL",
	"LOCALTIME",
	"LOCALTIMESTAMP",
	"LOCATION",
	"LOCK",
	"LOCKED",
	"LOGGED",
	"MAPPING",
	"MATCH",
	"MATERIALIZED",
	"MAXVALUE",
	"METHOD",
	"MINUTE",
	"MINVALUE",
	"MODE",
	"MONTH",
	"MOVE",
	"NAME",
	"NAMES",
	"NATIONAL",
	"NATURAL",
	"NCHAR",
	"NEW",
	"NEXT",
	"NFC",
	"NFD",
	"NFKC",
	"NFKD",
	"NO",
	"NONE",
	"NORMALIZE",
	"NORMALIZED",
	"NOT",
	"NOTHING",
	"NOTIFY",
	"NOTNULL",
	"NOWAIT",
	"NULL",
	"NULLIF",
	"NULLS",
	"NUMERIC",
	"OBJECT",
	"OF",
	"OFF",
	"OFFSET",
	"OIDS",
	"OLD",
	"ON",
	"ONLY",
	"OPERATOR",
	"OPTION",
	"OPTIONS",
	"OR",
	"ORDER",
	"ORDINALITY",
	"OTHERS",
	"OUT",
	"OUTER",
	"OVER",
	"OVERLAPS",
	"OVERLAY",
	"OVERRIDING",
	"OWNED",
	"OWNER",
	"PARALLEL",
	"PARSER",
	"PARTIAL",
	"PARTITION",
	"PASSING",
	"PASSWORD",
	"PLACING",
	"PLANS",
	"POLICY",
	"POSITION",
	"PRECEDING",
	"PRECISION",
	"PRESERVE",
	"PREPARE",
	"PREPARED",
	"PRIMARY",
	"PRIOR",
	"PRIVILEGES",
	"PROCEDURAL",
	"PROCEDURE",
	"PROCEDURES",
	"PROGRAM",
	"PUBLICATION",
	"QUOTE",
	"RANGE",
	"READ",
	"REAL",
	"REASSIGN",
	"RECHECK",
	"RECURSIVE",
	"REF",
	"REFERENCES",
	"REFERENCING",
	"REFRESH",
	"REINDEX",
	"RELATIVE",
	"RELEASE",
	"RENAME",
	"REPEATABLE",
	"REPLACE",
	"REPLICA",
	"RESET",
	"RESTART",
	"RESTRICT",
	"RETURNING",
	"RETURNS",
	"REVOKE",
	"RIGHT",
	"ROLE",
	"ROLLBACK",
	"ROLLUP",
	"ROUTINE",
	"ROUTINES",
	"ROW",
	"ROWS",
	"RULE",
	"SAVEPOINT",
	"SCHEMA",
	"SCHEMAS",
	"SCROLL",
	"SEARCH",
	"SECOND",
	"SECURITY",
	"SELECT",
	"SEQUENCE",
	"SEQUENCES",
	"SERIALIZABLE",
	"SERVER",
	"SESSION",
	"SESSION_USER",
	"SET",
	"SETS",
	"SETOF",
	"SHARE",
	"SHOW",
	"SIMILAR",
	"SIMPLE",
	"SKIP",
	"SMALLINT",
	"SNAPSHOT",
	"SOME",
	"SQL",
	"STABLE",
	"STANDALONE",
	"START",
	"STATEMENT",
	"STATISTICS",
	"STDIN",
	"STDOUT",
	"STORAGE",
	"STORED",
	"STRICT",
	"STRIP",
	"SUBSCRIPTION",
	"SUBSTRING",
	"SUPPORT",
	"SYMMETRIC",
	"SYSID",
	"SYSTEM",
	"TABLE",
	"TABLES",
	"TABLESAMPLE",
	"TABLESPACE",
	"TEMP",
	"TEMPLATE",
	"TEMPORARY",
	"TEXT",
	"THEN",
	"TIES",
	"TIME",
	"TIMESTAMP",
	"TO",
	"TRAILING",
	"TRANSACTION",
	"TRANSFORM",
	"TREAT",
	"TRIGGER",
	"TRIM",
	"TRUE",
	"TRUNCATE",
	"TRUSTED",
	"TYPE",
	"TYPES",
	"UESCAPE",
	"UNBOUNDED",
	"UNCOMMITTED",
	"UNENCRYPTED",
	"UNION",
	"UNIQUE",
	"UNKNOWN",
	"UNLISTEN",
	"UNLOGGED",
	"UNTIL",
	"UPDATE",
	"USER",
	"USING",
	"VACUUM",
	"VALID",
	"VALIDATE",
	"VALIDATOR",
	"VALUE",
	"VALUES",
	"VARCHAR",
	"VARIADIC",
	"VARYING",
	"VERBOSE",
	"VERSION",
	"VIEW",
	"VIEWS",
	"VOLATILE",
	"WHEN",
	"WHERE",
	"WHITESPACE",
	"WINDOW",
	"WITH",
	"WITHIN",
	"WITHOUT",
	"WORK",
	"WRAPPER",
	"WRITE",
	"XML",
	"XMLATTRIBUTES",
	"XMLCONCAT",
	"XMLELEMENT",
	"XMLEXISTS",
	"XMLFOREST",
	"XMLNAMESPACES",
	"XMLPARSE",
	"XMLPI",
	"XMLROOT",
	"XMLSERIALIZE",
	"XMLTABLE",
	"YEAR",
	"YES",
	"ZONE",
	"UNEXPECTED_SYMBOL",
	"Identifier",
	"Name",
	"';'",
	"','",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 366

var yyAct = [...]int{
	25, 10, 38, 13, 46, 18, 3, 43, 25, 48,
	32, 36, 37, 40, 23, 19, 21, 20, 11, 9,
	44, 2, 1, 14, 15, 16, 17, 22, 42, 0,
	0, 28, 29, 30, 4, 31, 0, 34, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 12, 45, 0,
	41, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 5, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 24, 0,
	0, 0, 0, 0, 0, 0, 24, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 6, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 26, 0, 0,
	0, 0, 0, 0, 0, 26, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 39, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 27, 33, 0, 0, 0, 0,
	0, 0, 27, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	7, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 47, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 35, 0,
	0, 0, 0, 0, 0, 8,
}

var yyPact = [...]int{
	2, -1000, -456, -389, -389, -389, -389, -389, -387, -1000,
	-1000, -1, -1000, -1000, -92, -1, -1, -1, -92, -1000,
	-38, -1000, -100, -1000, -199, -1000, -88, -261, -1000, -1000,
	-1000, -1000, -1000, -35, -1000, -92, -297, -1000, -1000, -1000,
	-1000, -1000, -1000, -61, -295, -1000, -1000, -1000, -1000,
}

var yyPgo = [...]int{
	0, 15, 28, 14, 27, 16, 22, 21, 19, 18,
}

var yyR1 = [...]int{
	0, 6, 8, 8, 1, 1, 1, 2, 2, 2,
	2, 9, 9, 9, 3, 3, 3, 3, 3, 4,
	4, 4, 5, 5, 7, 7, 7, 7, 7, 7,
}

var yyR2 = [...]int{
	0, 2, 0, 1, 0, 2, 3, 2, 2, 2,
	1, 0, 1, 1, 3, 1, 2, 2, 2, 1,
	2, 3, 0, 1, 3, 3, 3, 3, 3, 3,
}

var yyChk = [...]int{
	-1000, -6, -7, 4, 32, 64, 123, 328, 363, -8,
	457, -9, 436, 392, -9, -9, -9, -9, 392, -1,
	18, -5, -4, -3, 198, 100, 247, 304, -1, -1,
	-1, -5, 48, 243, -3, 458, 210, 100, 263, 438,
	48, -3, -2, 304, 317, 345, 65, 404, 304,
}

var yyDef = [...]int{
	0, -2, 2, 11, 11, 11, 11, 11, 0, 1,
	3, 4, 12, 13, 22, 4, 4, 4, 22, 24,
	0, 25, 23, 19, 0, 15, 0, 0, 26, 27,
	28, 29, 5, 0, 20, 0, 0, 16, 17, 18,
	6, 21, 14, 0, 0, 10, 7, 8, 9,
}

var yyTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 458, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 457,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 90, 91,
	92, 93, 94, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 105, 106, 107, 108, 109, 110, 111,
	112, 113, 114, 115, 116, 117, 118, 119, 120, 121,
	122, 123, 124, 125, 126, 127, 128, 129, 130, 131,
	132, 133, 134, 135, 136, 137, 138, 139, 140, 141,
	142, 143, 144, 145, 146, 147, 148, 149, 150, 151,
	152, 153, 154, 155, 156, 157, 158, 159, 160, 161,
	162, 163, 164, 165, 166, 167, 168, 169, 170, 171,
	172, 173, 174, 175, 176, 177, 178, 179, 180, 181,
	182, 183, 184, 185, 186, 187, 188, 189, 190, 191,
	192, 193, 194, 195, 196, 197, 198, 199, 200, 201,
	202, 203, 204, 205, 206, 207, 208, 209, 210, 211,
	212, 213, 214, 215, 216, 217, 218, 219, 220, 221,
	222, 223, 224, 225, 226, 227, 228, 229, 230, 231,
	232, 233, 234, 235, 236, 237, 238, 239, 240, 241,
	242, 243, 244, 245, 246, 247, 248, 249, 250, 251,
	252, 253, 254, 255, 256, 257,
}

var yyTok3 = [...]int{
	57600, 258, 57601, 259, 57602, 260, 57603, 261, 57604, 262,
	57605, 263, 57606, 264, 57607, 265, 57608, 266, 57609, 267,
	57610, 268, 57611, 269, 57612, 270, 57613, 271, 57614, 272,
	57615, 273, 57616, 274, 57617, 275, 57618, 276, 57619, 277,
	57620, 278, 57621, 279, 57622, 280, 57623, 281, 57624, 282,
	57625, 283, 57626, 284, 57627, 285, 57628, 286, 57629, 287,
	57630, 288, 57631, 289, 57632, 290, 57633, 291, 57634, 292,
	57635, 293, 57636, 294, 57637, 295, 57638, 296, 57639, 297,
	57640, 298, 57641, 299, 57642, 300, 57643, 301, 57644, 302,
	57645, 303, 57646, 304, 57647, 305, 57648, 306, 57649, 307,
	57650, 308, 57651, 309, 57652, 310, 57653, 311, 57654, 312,
	57655, 313, 57656, 314, 57657, 315, 57658, 316, 57659, 317,
	57660, 318, 57661, 319, 57662, 320, 57663, 321, 57664, 322,
	57665, 323, 57666, 324, 57667, 325, 57668, 326, 57669, 327,
	57670, 328, 57671, 329, 57672, 330, 57673, 331, 57674, 332,
	57675, 333, 57676, 334, 57677, 335, 57678, 336, 57679, 337,
	57680, 338, 57681, 339, 57682, 340, 57683, 341, 57684, 342,
	57685, 343, 57686, 344, 57687, 345, 57688, 346, 57689, 347,
	57690, 348, 57691, 349, 57692, 350, 57693, 351, 57694, 352,
	57695, 353, 57696, 354, 57697, 355, 57698, 356, 57699, 357,
	57700, 358, 57701, 359, 57702, 360, 57703, 361, 57704, 362,
	57705, 363, 57706, 364, 57707, 365, 57708, 366, 57709, 367,
	57710, 368, 57711, 369, 57712, 370, 57713, 371, 57714, 372,
	57715, 373, 57716, 374, 57717, 375, 57718, 376, 57719, 377,
	57720, 378, 57721, 379, 57722, 380, 57723, 381, 57724, 382,
	57725, 383, 57726, 384, 57727, 385, 57728, 386, 57729, 387,
	57730, 388, 57731, 389, 57732, 390, 57733, 391, 57734, 392,
	57735, 393, 57736, 394, 57737, 395, 57738, 396, 57739, 397,
	57740, 398, 57741, 399, 57742, 400, 57743, 401, 57744, 402,
	57745, 403, 57746, 404, 57747, 405, 57748, 406, 57749, 407,
	57750, 408, 57751, 409, 57752, 410, 57753, 411, 57754, 412,
	57755, 413, 57756, 414, 57757, 415, 57758, 416, 57759, 417,
	57760, 418, 57761, 419, 57762, 420, 57763, 421, 57764, 422,
	57765, 423, 57766, 424, 57767, 425, 57768, 426, 57769, 427,
	57770, 428, 57771, 429, 57772, 430, 57773, 431, 57774, 432,
	57775, 433, 57776, 434, 57777, 435, 57778, 436, 57779, 437,
	57780, 438, 57781, 439, 57782, 440, 57783, 441, 57784, 442,
	57785, 443, 57786, 444, 57787, 445, 57788, 446, 57789, 447,
	57790, 448, 57791, 449, 57792, 450, 57793, 451, 57794, 452,
	57795, 453, 57796, 454, 57797, 455, 57798, 456, 0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:160
		{
			yyVAL.bool = false
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:161
		{
			yyVAL.bool = true
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:162
		{
			yyVAL.bool = false
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:165
		{
			yyVAL.str = "read committed"
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:166
		{
			yyVAL.str = "read uncommitted"
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:167
		{
			yyVAL.str = "repeatable read"
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:168
		{
			yyVAL.str = "serializable"
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:176
		{
			yyVAL.opt = newOption("isolation_level", yyDollar[3].str)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:179
		{
			yyVAL.opt = newOption("deferrable", true)
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:180
		{
			yyVAL.opt = newOption("deferrable", false)
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:181
		{
			yyVAL.opt = newOption("read_only", true)
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:182
		{
			yyVAL.opt = newOption("read_only", false)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:185
		{
			yyVAL.opts = newOptionList(yyDollar[1].opt)
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammar.y:186
		{
			yyVAL.opts = append(yyDollar[1].opts, yyDollar[2].opt)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:187
		{
			yyVAL.opts = append(yyDollar[1].opts, yyDollar[3].opt)
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammar.y:190
		{
			yyVAL.opts = nil
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammar.y:191
		{
			yyVAL.opts = yyDollar[1].opts
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:194
		{
			stmt := &sql.RollbackStmt{Chain: yyDollar[3].bool}
			yylex.(*Lexer).Statement = stmt
			if yyDebug > 6 {
				__yyfmt__.Printf("stmt = %#v\n", stmt)
			}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:201
		{
			stmt := newBeginStmt(yyDollar[3].opts)
			yylex.(*Lexer).Statement = stmt
			if yyDebug > 6 {
				__yyfmt__.Printf("stmt = %#v\n", stmt)
			}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:208
		{
			stmt := &sql.CommitStmt{Chain: yyDollar[3].bool}
			yylex.(*Lexer).Statement = stmt
			if yyDebug > 6 {
				__yyfmt__.Printf("stmt = %#v\n", stmt)
			}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:215
		{
			stmt := &sql.CommitStmt{Chain: yyDollar[3].bool}
			yylex.(*Lexer).Statement = stmt
			if yyDebug > 6 {
				__yyfmt__.Printf("stmt = %#v\n", stmt)
			}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:222
		{
			stmt := &sql.RollbackStmt{Chain: yyDollar[3].bool}
			yylex.(*Lexer).Statement = stmt
			if yyDebug > 6 {
				__yyfmt__.Printf("stmt = %#v\n", stmt)
			}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammar.y:229
		{
			stmt := newBeginStmt(yyDollar[3].opts)
			yylex.(*Lexer).Statement = stmt
			if yyDebug > 6 {
				__yyfmt__.Printf("stmt = %#v\n", stmt)
			}
		}
	}
	goto yystack /* stack new state and value */
}
