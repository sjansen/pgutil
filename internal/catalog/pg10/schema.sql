-- PostgreSQL 10.18 on x86_64-apple-darwin14.5.0, compiled by Apple LLVM version 7.0.2 (clang-700.1.81), 64-bit
CREATE TABLE pg_aggregate (
    aggfnoid         regproc  NOT NULL
,   aggkind          "char"   NOT NULL
,   aggnumdirectargs smallint NOT NULL
,   aggtransfn       regproc  NOT NULL
,   aggfinalfn       regproc  NOT NULL
,   aggcombinefn     regproc  NOT NULL
,   aggserialfn      regproc  NOT NULL
,   aggdeserialfn    regproc  NOT NULL
,   aggmtransfn      regproc  NOT NULL
,   aggminvtransfn   regproc  NOT NULL
,   aggmfinalfn      regproc  NOT NULL
,   aggfinalextra    boolean  NOT NULL
,   aggmfinalextra   boolean  NOT NULL
,   aggsortop        oid      NOT NULL
,   aggtranstype     oid      NOT NULL
,   aggtransspace    integer  NOT NULL
,   aggmtranstype    oid      NOT NULL
,   aggmtransspace   integer  NOT NULL
,   agginitval       text
,   aggminitval      text
)
;
CREATE TABLE pg_am (
    amname    name    NOT NULL
,   amhandler regproc NOT NULL
,   amtype    "char"  NOT NULL
)
;
CREATE TABLE pg_amop (
    amopfamily     oid      NOT NULL
,   amoplefttype   oid      NOT NULL
,   amoprighttype  oid      NOT NULL
,   amopstrategy   smallint NOT NULL
,   amoppurpose    "char"   NOT NULL
,   amopopr        oid      NOT NULL
,   amopmethod     oid      NOT NULL
,   amopsortfamily oid      NOT NULL
)
;
CREATE TABLE pg_amproc (
    amprocfamily    oid      NOT NULL
,   amproclefttype  oid      NOT NULL
,   amprocrighttype oid      NOT NULL
,   amprocnum       smallint NOT NULL
,   amproc          regproc  NOT NULL
)
;
CREATE TABLE pg_attrdef (
    adrelid oid          NOT NULL
,   adnum   smallint     NOT NULL
,   adbin   pg_node_tree
,   adsrc   text
)
;
CREATE TABLE pg_attribute (
    attrelid      oid       NOT NULL
,   attname       name      NOT NULL
,   atttypid      oid       NOT NULL
,   attstattarget integer   NOT NULL
,   attlen        smallint  NOT NULL
,   attnum        smallint  NOT NULL
,   attndims      integer   NOT NULL
,   attcacheoff   integer   NOT NULL
,   atttypmod     integer   NOT NULL
,   attbyval      boolean   NOT NULL
,   attstorage    "char"    NOT NULL
,   attalign      "char"    NOT NULL
,   attnotnull    boolean   NOT NULL
,   atthasdef     boolean   NOT NULL
,   attidentity   "char"    NOT NULL
,   attisdropped  boolean   NOT NULL
,   attislocal    boolean   NOT NULL
,   attinhcount   integer   NOT NULL
,   attcollation  oid       NOT NULL
,   attacl        aclitem[]
,   attoptions    text[]
,   attfdwoptions text[]
)
;
CREATE TABLE pg_auth_members (
    roleid       oid     NOT NULL
,   member       oid     NOT NULL
,   grantor      oid     NOT NULL
,   admin_option boolean NOT NULL
)
;
CREATE TABLE pg_authid (
    rolname        name                     NOT NULL
,   rolsuper       boolean                  NOT NULL
,   rolinherit     boolean                  NOT NULL
,   rolcreaterole  boolean                  NOT NULL
,   rolcreatedb    boolean                  NOT NULL
,   rolcanlogin    boolean                  NOT NULL
,   rolreplication boolean                  NOT NULL
,   rolbypassrls   boolean                  NOT NULL
,   rolconnlimit   integer                  NOT NULL
,   rolpassword    text
,   rolvaliduntil  timestamp with time zone
)
;
CREATE TABLE pg_available_extension_versions (
    name        name
,   version     text
,   installed   boolean
,   superuser   boolean
,   relocatable boolean
,   schema      name
,   requires    name[]
,   comment     text
)
;
CREATE TABLE pg_available_extensions (
    name              name
,   default_version   text
,   installed_version text
,   comment           text
)
;
CREATE TABLE pg_cast (
    castsource  oid    NOT NULL
,   casttarget  oid    NOT NULL
,   castfunc    oid    NOT NULL
,   castcontext "char" NOT NULL
,   castmethod  "char" NOT NULL
)
;
CREATE TABLE pg_class (
    relname             name         NOT NULL
,   relnamespace        oid          NOT NULL
,   reltype             oid          NOT NULL
,   reloftype           oid          NOT NULL
,   relowner            oid          NOT NULL
,   relam               oid          NOT NULL
,   relfilenode         oid          NOT NULL
,   reltablespace       oid          NOT NULL
,   relpages            integer      NOT NULL
,   reltuples           real         NOT NULL
,   relallvisible       integer      NOT NULL
,   reltoastrelid       oid          NOT NULL
,   relhasindex         boolean      NOT NULL
,   relisshared         boolean      NOT NULL
,   relpersistence      "char"       NOT NULL
,   relkind             "char"       NOT NULL
,   relnatts            smallint     NOT NULL
,   relchecks           smallint     NOT NULL
,   relhasoids          boolean      NOT NULL
,   relhaspkey          boolean      NOT NULL
,   relhasrules         boolean      NOT NULL
,   relhastriggers      boolean      NOT NULL
,   relhassubclass      boolean      NOT NULL
,   relrowsecurity      boolean      NOT NULL
,   relforcerowsecurity boolean      NOT NULL
,   relispopulated      boolean      NOT NULL
,   relreplident        "char"       NOT NULL
,   relispartition      boolean      NOT NULL
,   relfrozenxid        xid          NOT NULL
,   relminmxid          xid          NOT NULL
,   relacl              aclitem[]
,   reloptions          text[]
,   relpartbound        pg_node_tree
)
;
CREATE TABLE pg_collation (
    collname      name    NOT NULL
,   collnamespace oid     NOT NULL
,   collowner     oid     NOT NULL
,   collprovider  "char"  NOT NULL
,   collencoding  integer NOT NULL
,   collcollate   name    NOT NULL
,   collctype     name    NOT NULL
,   collversion   text
)
;
CREATE TABLE pg_config (
    name    text
,   setting text
)
;
CREATE TABLE pg_constraint (
    conname       name         NOT NULL
,   connamespace  oid          NOT NULL
,   contype       "char"       NOT NULL
,   condeferrable boolean      NOT NULL
,   condeferred   boolean      NOT NULL
,   convalidated  boolean      NOT NULL
,   conrelid      oid          NOT NULL
,   contypid      oid          NOT NULL
,   conindid      oid          NOT NULL
,   confrelid     oid          NOT NULL
,   confupdtype   "char"       NOT NULL
,   confdeltype   "char"       NOT NULL
,   confmatchtype "char"       NOT NULL
,   conislocal    boolean      NOT NULL
,   coninhcount   integer      NOT NULL
,   connoinherit  boolean      NOT NULL
,   conkey        smallint[]
,   confkey       smallint[]
,   conpfeqop     oid[]
,   conppeqop     oid[]
,   conffeqop     oid[]
,   conexclop     oid[]
,   conbin        pg_node_tree
,   consrc        text
)
;
CREATE TABLE pg_conversion (
    conname        name    NOT NULL
,   connamespace   oid     NOT NULL
,   conowner       oid     NOT NULL
,   conforencoding integer NOT NULL
,   contoencoding  integer NOT NULL
,   conproc        regproc NOT NULL
,   condefault     boolean NOT NULL
)
;
CREATE TABLE pg_cursors (
    name          text
,   statement     text
,   is_holdable   boolean
,   is_binary     boolean
,   is_scrollable boolean
,   creation_time timestamp with time zone
)
;
CREATE TABLE pg_database (
    datname       name      NOT NULL
,   datdba        oid       NOT NULL
,   encoding      integer   NOT NULL
,   datcollate    name      NOT NULL
,   datctype      name      NOT NULL
,   datistemplate boolean   NOT NULL
,   datallowconn  boolean   NOT NULL
,   datconnlimit  integer   NOT NULL
,   datlastsysoid oid       NOT NULL
,   datfrozenxid  xid       NOT NULL
,   datminmxid    xid       NOT NULL
,   dattablespace oid       NOT NULL
,   datacl        aclitem[]
)
;
CREATE TABLE pg_db_role_setting (
    setdatabase oid    NOT NULL
,   setrole     oid    NOT NULL
,   setconfig   text[]
)
;
CREATE TABLE pg_default_acl (
    defaclrole      oid       NOT NULL
,   defaclnamespace oid       NOT NULL
,   defaclobjtype   "char"    NOT NULL
,   defaclacl       aclitem[]
)
;
CREATE TABLE pg_depend (
    classid     oid     NOT NULL
,   objid       oid     NOT NULL
,   objsubid    integer NOT NULL
,   refclassid  oid     NOT NULL
,   refobjid    oid     NOT NULL
,   refobjsubid integer NOT NULL
,   deptype     "char"  NOT NULL
)
;
CREATE TABLE pg_description (
    objoid      oid     NOT NULL
,   classoid    oid     NOT NULL
,   objsubid    integer NOT NULL
,   description text    NOT NULL
)
;
CREATE TABLE pg_enum (
    enumtypid     oid  NOT NULL
,   enumsortorder real NOT NULL
,   enumlabel     name NOT NULL
)
;
CREATE TABLE pg_event_trigger (
    evtname    name   NOT NULL
,   evtevent   name   NOT NULL
,   evtowner   oid    NOT NULL
,   evtfoid    oid    NOT NULL
,   evtenabled "char" NOT NULL
,   evttags    text[]
)
;
CREATE TABLE pg_extension (
    extname        name    NOT NULL
,   extowner       oid     NOT NULL
,   extnamespace   oid     NOT NULL
,   extrelocatable boolean NOT NULL
,   extversion     text    NOT NULL
,   extconfig      oid[]
,   extcondition   text[]
)
;
CREATE TABLE pg_file_settings (
    sourcefile text
,   sourceline integer
,   seqno      integer
,   name       text
,   setting    text
,   applied    boolean
,   error      text
)
;
CREATE TABLE pg_foreign_data_wrapper (
    fdwname      name      NOT NULL
,   fdwowner     oid       NOT NULL
,   fdwhandler   oid       NOT NULL
,   fdwvalidator oid       NOT NULL
,   fdwacl       aclitem[]
,   fdwoptions   text[]
)
;
CREATE TABLE pg_foreign_server (
    srvname    name      NOT NULL
,   srvowner   oid       NOT NULL
,   srvfdw     oid       NOT NULL
,   srvtype    text
,   srvversion text
,   srvacl     aclitem[]
,   srvoptions text[]
)
;
CREATE TABLE pg_foreign_table (
    ftrelid   oid    NOT NULL
,   ftserver  oid    NOT NULL
,   ftoptions text[]
)
;
CREATE TABLE pg_group (
    groname  name
,   grosysid oid
,   grolist  oid[]
)
;
CREATE TABLE pg_hba_file_rules (
    line_number integer
,   type        text
,   database    text[]
,   user_name   text[]
,   address     text
,   netmask     text
,   auth_method text
,   options     text[]
,   error       text
)
;
CREATE TABLE pg_index (
    indexrelid     oid          NOT NULL
,   indrelid       oid          NOT NULL
,   indnatts       smallint     NOT NULL
,   indisunique    boolean      NOT NULL
,   indisprimary   boolean      NOT NULL
,   indisexclusion boolean      NOT NULL
,   indimmediate   boolean      NOT NULL
,   indisclustered boolean      NOT NULL
,   indisvalid     boolean      NOT NULL
,   indcheckxmin   boolean      NOT NULL
,   indisready     boolean      NOT NULL
,   indislive      boolean      NOT NULL
,   indisreplident boolean      NOT NULL
,   indkey         smallint[]   NOT NULL
,   indcollation   oid[]        NOT NULL
,   indclass       oid[]        NOT NULL
,   indoption      smallint[]   NOT NULL
,   indexprs       pg_node_tree
,   indpred        pg_node_tree
)
;
CREATE TABLE pg_indexes (
    schemaname name
,   tablename  name
,   indexname  name
,   tablespace name
,   indexdef   text
)
;
CREATE TABLE pg_inherits (
    inhrelid  oid     NOT NULL
,   inhparent oid     NOT NULL
,   inhseqno  integer NOT NULL
)
;
CREATE TABLE pg_init_privs (
    objoid    oid       NOT NULL
,   classoid  oid       NOT NULL
,   objsubid  integer   NOT NULL
,   privtype  "char"    NOT NULL
,   initprivs aclitem[] NOT NULL
)
;
CREATE TABLE pg_language (
    lanname       name      NOT NULL
,   lanowner      oid       NOT NULL
,   lanispl       boolean   NOT NULL
,   lanpltrusted  boolean   NOT NULL
,   lanplcallfoid oid       NOT NULL
,   laninline     oid       NOT NULL
,   lanvalidator  oid       NOT NULL
,   lanacl        aclitem[]
)
;
CREATE TABLE pg_largeobject (
    loid   oid     NOT NULL
,   pageno integer NOT NULL
,   data   bytea   NOT NULL
)
;
CREATE TABLE pg_largeobject_metadata (
    lomowner oid       NOT NULL
,   lomacl   aclitem[]
)
;
CREATE TABLE pg_locks (
    locktype           text
,   database           oid
,   relation           oid
,   page               integer
,   tuple              smallint
,   virtualxid         text
,   transactionid      xid
,   classid            oid
,   objid              oid
,   objsubid           smallint
,   virtualtransaction text
,   pid                integer
,   mode               text
,   granted            boolean
,   fastpath           boolean
)
;
CREATE TABLE pg_matviews (
    schemaname   name
,   matviewname  name
,   matviewowner name
,   tablespace   name
,   hasindexes   boolean
,   ispopulated  boolean
,   definition   text
)
;
CREATE TABLE pg_namespace (
    nspname  name      NOT NULL
,   nspowner oid       NOT NULL
,   nspacl   aclitem[]
)
;
CREATE TABLE pg_opclass (
    opcmethod    oid     NOT NULL
,   opcname      name    NOT NULL
,   opcnamespace oid     NOT NULL
,   opcowner     oid     NOT NULL
,   opcfamily    oid     NOT NULL
,   opcintype    oid     NOT NULL
,   opcdefault   boolean NOT NULL
,   opckeytype   oid     NOT NULL
)
;
CREATE TABLE pg_operator (
    oprname      name    NOT NULL
,   oprnamespace oid     NOT NULL
,   oprowner     oid     NOT NULL
,   oprkind      "char"  NOT NULL
,   oprcanmerge  boolean NOT NULL
,   oprcanhash   boolean NOT NULL
,   oprleft      oid     NOT NULL
,   oprright     oid     NOT NULL
,   oprresult    oid     NOT NULL
,   oprcom       oid     NOT NULL
,   oprnegate    oid     NOT NULL
,   oprcode      regproc NOT NULL
,   oprrest      regproc NOT NULL
,   oprjoin      regproc NOT NULL
)
;
CREATE TABLE pg_opfamily (
    opfmethod    oid  NOT NULL
,   opfname      name NOT NULL
,   opfnamespace oid  NOT NULL
,   opfowner     oid  NOT NULL
)
;
CREATE TABLE pg_partitioned_table (
    partrelid     oid          NOT NULL
,   partstrat     "char"       NOT NULL
,   partnatts     smallint     NOT NULL
,   partattrs     smallint[]   NOT NULL
,   partclass     oid[]        NOT NULL
,   partcollation oid[]        NOT NULL
,   partexprs     pg_node_tree
)
;
CREATE TABLE pg_pltemplate (
    tmplname      name      NOT NULL
,   tmpltrusted   boolean   NOT NULL
,   tmpldbacreate boolean   NOT NULL
,   tmplhandler   text      NOT NULL
,   tmplinline    text
,   tmplvalidator text
,   tmpllibrary   text      NOT NULL
,   tmplacl       aclitem[]
)
;
CREATE TABLE pg_policies (
    schemaname name
,   tablename  name
,   policyname name
,   permissive text
,   roles      name[]
,   cmd        text
,   qual       text
,   with_check text
)
;
CREATE TABLE pg_policy (
    polname       name         NOT NULL
,   polrelid      oid          NOT NULL
,   polcmd        "char"       NOT NULL
,   polpermissive boolean      NOT NULL
,   polroles      oid[]
,   polqual       pg_node_tree
,   polwithcheck  pg_node_tree
)
;
CREATE TABLE pg_prepared_statements (
    name            text
,   statement       text
,   prepare_time    timestamp with time zone
,   parameter_types regtype[]
,   from_sql        boolean
)
;
CREATE TABLE pg_prepared_xacts (
    transaction xid
,   gid         text
,   prepared    timestamp with time zone
,   owner       name
,   database    name
)
;
CREATE TABLE pg_proc (
    proname         name         NOT NULL
,   pronamespace    oid          NOT NULL
,   proowner        oid          NOT NULL
,   prolang         oid          NOT NULL
,   procost         real         NOT NULL
,   prorows         real         NOT NULL
,   provariadic     oid          NOT NULL
,   protransform    regproc      NOT NULL
,   proisagg        boolean      NOT NULL
,   proiswindow     boolean      NOT NULL
,   prosecdef       boolean      NOT NULL
,   proleakproof    boolean      NOT NULL
,   proisstrict     boolean      NOT NULL
,   proretset       boolean      NOT NULL
,   provolatile     "char"       NOT NULL
,   proparallel     "char"       NOT NULL
,   pronargs        smallint     NOT NULL
,   pronargdefaults smallint     NOT NULL
,   prorettype      oid          NOT NULL
,   proargtypes     oid[]        NOT NULL
,   proallargtypes  oid[]
,   proargmodes     "char"[]
,   proargnames     text[]
,   proargdefaults  pg_node_tree
,   protrftypes     oid[]
,   prosrc          text         NOT NULL
,   probin          text
,   proconfig       text[]
,   proacl          aclitem[]
)
;
CREATE TABLE pg_publication (
    pubname      name    NOT NULL
,   pubowner     oid     NOT NULL
,   puballtables boolean NOT NULL
,   pubinsert    boolean NOT NULL
,   pubupdate    boolean NOT NULL
,   pubdelete    boolean NOT NULL
)
;
CREATE TABLE pg_publication_rel (
    prpubid oid NOT NULL
,   prrelid oid NOT NULL
)
;
CREATE TABLE pg_publication_tables (
    pubname    name
,   schemaname name
,   tablename  name
)
;
CREATE TABLE pg_range (
    rngtypid     oid     NOT NULL
,   rngsubtype   oid     NOT NULL
,   rngcollation oid     NOT NULL
,   rngsubopc    oid     NOT NULL
,   rngcanonical regproc NOT NULL
,   rngsubdiff   regproc NOT NULL
)
;
CREATE TABLE pg_replication_origin (
    roident oid  NOT NULL
,   roname  text NOT NULL
)
;
CREATE TABLE pg_replication_origin_status (
    local_id    oid
,   external_id text
,   remote_lsn  pg_lsn
,   local_lsn   pg_lsn
)
;
CREATE TABLE pg_replication_slots (
    slot_name           name
,   plugin              name
,   slot_type           text
,   datoid              oid
,   database            name
,   temporary           boolean
,   active              boolean
,   active_pid          integer
,   xmin                xid
,   catalog_xmin        xid
,   restart_lsn         pg_lsn
,   confirmed_flush_lsn pg_lsn
)
;
CREATE TABLE pg_rewrite (
    rulename   name         NOT NULL
,   ev_class   oid          NOT NULL
,   ev_type    "char"       NOT NULL
,   ev_enabled "char"       NOT NULL
,   is_instead boolean      NOT NULL
,   ev_qual    pg_node_tree
,   ev_action  pg_node_tree
)
;
CREATE TABLE pg_roles (
    rolname        name
,   rolsuper       boolean
,   rolinherit     boolean
,   rolcreaterole  boolean
,   rolcreatedb    boolean
,   rolcanlogin    boolean
,   rolreplication boolean
,   rolconnlimit   integer
,   rolpassword    text
,   rolvaliduntil  timestamp with time zone
,   rolbypassrls   boolean
,   rolconfig      text[]
,   oid            oid
)
;
CREATE TABLE pg_rules (
    schemaname name
,   tablename  name
,   rulename   name
,   definition text
)
;
CREATE TABLE pg_seclabel (
    objoid   oid     NOT NULL
,   classoid oid     NOT NULL
,   objsubid integer NOT NULL
,   provider text    NOT NULL
,   label    text    NOT NULL
)
;
CREATE TABLE pg_seclabels (
    objoid       oid
,   classoid     oid
,   objsubid     integer
,   objtype      text
,   objnamespace oid
,   objname      text
,   provider     text
,   label        text
)
;
CREATE TABLE pg_sequence (
    seqrelid     oid     NOT NULL
,   seqtypid     oid     NOT NULL
,   seqstart     bigint  NOT NULL
,   seqincrement bigint  NOT NULL
,   seqmax       bigint  NOT NULL
,   seqmin       bigint  NOT NULL
,   seqcache     bigint  NOT NULL
,   seqcycle     boolean NOT NULL
)
;
CREATE TABLE pg_sequences (
    schemaname    name
,   sequencename  name
,   sequenceowner name
,   data_type     regtype
,   start_value   bigint
,   min_value     bigint
,   max_value     bigint
,   increment_by  bigint
,   cycle         boolean
,   cache_size    bigint
,   last_value    bigint
)
;
CREATE TABLE pg_settings (
    name            text
,   setting         text
,   unit            text
,   category        text
,   short_desc      text
,   extra_desc      text
,   context         text
,   vartype         text
,   source          text
,   min_val         text
,   max_val         text
,   enumvals        text[]
,   boot_val        text
,   reset_val       text
,   sourcefile      text
,   sourceline      integer
,   pending_restart boolean
)
;
CREATE TABLE pg_shadow (
    usename      name
,   usesysid     oid
,   usecreatedb  boolean
,   usesuper     boolean
,   userepl      boolean
,   usebypassrls boolean
,   passwd       text
,   valuntil     abstime
,   useconfig    text[]
)
;
CREATE TABLE pg_shdepend (
    dbid       oid     NOT NULL
,   classid    oid     NOT NULL
,   objid      oid     NOT NULL
,   objsubid   integer NOT NULL
,   refclassid oid     NOT NULL
,   refobjid   oid     NOT NULL
,   deptype    "char"  NOT NULL
)
;
CREATE TABLE pg_shdescription (
    objoid      oid  NOT NULL
,   classoid    oid  NOT NULL
,   description text NOT NULL
)
;
CREATE TABLE pg_shseclabel (
    objoid   oid  NOT NULL
,   classoid oid  NOT NULL
,   provider text NOT NULL
,   label    text NOT NULL
)
;
CREATE TABLE pg_stat_activity (
    datid            oid
,   datname          name
,   pid              integer
,   usesysid         oid
,   usename          name
,   application_name text
,   client_addr      inet
,   client_hostname  text
,   client_port      integer
,   backend_start    timestamp with time zone
,   xact_start       timestamp with time zone
,   query_start      timestamp with time zone
,   state_change     timestamp with time zone
,   wait_event_type  text
,   wait_event       text
,   state            text
,   backend_xid      xid
,   backend_xmin     xid
,   query            text
,   backend_type     text
)
;
CREATE TABLE pg_stat_all_indexes (
    relid         oid
,   indexrelid    oid
,   schemaname    name
,   relname       name
,   indexrelname  name
,   idx_scan      bigint
,   idx_tup_read  bigint
,   idx_tup_fetch bigint
)
;
CREATE TABLE pg_stat_all_tables (
    relid               oid
,   schemaname          name
,   relname             name
,   seq_scan            bigint
,   seq_tup_read        bigint
,   idx_scan            bigint
,   idx_tup_fetch       bigint
,   n_tup_ins           bigint
,   n_tup_upd           bigint
,   n_tup_del           bigint
,   n_tup_hot_upd       bigint
,   n_live_tup          bigint
,   n_dead_tup          bigint
,   n_mod_since_analyze bigint
,   last_vacuum         timestamp with time zone
,   last_autovacuum     timestamp with time zone
,   last_analyze        timestamp with time zone
,   last_autoanalyze    timestamp with time zone
,   vacuum_count        bigint
,   autovacuum_count    bigint
,   analyze_count       bigint
,   autoanalyze_count   bigint
)
;
CREATE TABLE pg_stat_archiver (
    archived_count     bigint
,   last_archived_wal  text
,   last_archived_time timestamp with time zone
,   failed_count       bigint
,   last_failed_wal    text
,   last_failed_time   timestamp with time zone
,   stats_reset        timestamp with time zone
)
;
CREATE TABLE pg_stat_bgwriter (
    checkpoints_timed     bigint
,   checkpoints_req       bigint
,   checkpoint_write_time double precision
,   checkpoint_sync_time  double precision
,   buffers_checkpoint    bigint
,   buffers_clean         bigint
,   maxwritten_clean      bigint
,   buffers_backend       bigint
,   buffers_backend_fsync bigint
,   buffers_alloc         bigint
,   stats_reset           timestamp with time zone
)
;
CREATE TABLE pg_stat_database (
    datid          oid
,   datname        name
,   numbackends    integer
,   xact_commit    bigint
,   xact_rollback  bigint
,   blks_read      bigint
,   blks_hit       bigint
,   tup_returned   bigint
,   tup_fetched    bigint
,   tup_inserted   bigint
,   tup_updated    bigint
,   tup_deleted    bigint
,   conflicts      bigint
,   temp_files     bigint
,   temp_bytes     bigint
,   deadlocks      bigint
,   blk_read_time  double precision
,   blk_write_time double precision
,   stats_reset    timestamp with time zone
)
;
CREATE TABLE pg_stat_database_conflicts (
    datid            oid
,   datname          name
,   confl_tablespace bigint
,   confl_lock       bigint
,   confl_snapshot   bigint
,   confl_bufferpin  bigint
,   confl_deadlock   bigint
)
;
CREATE TABLE pg_stat_progress_vacuum (
    pid                integer
,   datid              oid
,   datname            name
,   relid              oid
,   phase              text
,   heap_blks_total    bigint
,   heap_blks_scanned  bigint
,   heap_blks_vacuumed bigint
,   index_vacuum_count bigint
,   max_dead_tuples    bigint
,   num_dead_tuples    bigint
)
;
CREATE TABLE pg_stat_replication (
    pid              integer
,   usesysid         oid
,   usename          name
,   application_name text
,   client_addr      inet
,   client_hostname  text
,   client_port      integer
,   backend_start    timestamp with time zone
,   backend_xmin     xid
,   state            text
,   sent_lsn         pg_lsn
,   write_lsn        pg_lsn
,   flush_lsn        pg_lsn
,   replay_lsn       pg_lsn
,   write_lag        interval
,   flush_lag        interval
,   replay_lag       interval
,   sync_priority    integer
,   sync_state       text
)
;
CREATE TABLE pg_stat_ssl (
    pid         integer
,   ssl         boolean
,   version     text
,   cipher      text
,   bits        integer
,   compression boolean
,   clientdn    text
)
;
CREATE TABLE pg_stat_subscription (
    subid                 oid
,   subname               name
,   pid                   integer
,   relid                 oid
,   received_lsn          pg_lsn
,   last_msg_send_time    timestamp with time zone
,   last_msg_receipt_time timestamp with time zone
,   latest_end_lsn        pg_lsn
,   latest_end_time       timestamp with time zone
)
;
CREATE TABLE pg_stat_sys_indexes (
    relid         oid
,   indexrelid    oid
,   schemaname    name
,   relname       name
,   indexrelname  name
,   idx_scan      bigint
,   idx_tup_read  bigint
,   idx_tup_fetch bigint
)
;
CREATE TABLE pg_stat_sys_tables (
    relid               oid
,   schemaname          name
,   relname             name
,   seq_scan            bigint
,   seq_tup_read        bigint
,   idx_scan            bigint
,   idx_tup_fetch       bigint
,   n_tup_ins           bigint
,   n_tup_upd           bigint
,   n_tup_del           bigint
,   n_tup_hot_upd       bigint
,   n_live_tup          bigint
,   n_dead_tup          bigint
,   n_mod_since_analyze bigint
,   last_vacuum         timestamp with time zone
,   last_autovacuum     timestamp with time zone
,   last_analyze        timestamp with time zone
,   last_autoanalyze    timestamp with time zone
,   vacuum_count        bigint
,   autovacuum_count    bigint
,   analyze_count       bigint
,   autoanalyze_count   bigint
)
;
CREATE TABLE pg_stat_user_functions (
    funcid     oid
,   schemaname name
,   funcname   name
,   calls      bigint
,   total_time double precision
,   self_time  double precision
)
;
CREATE TABLE pg_stat_user_indexes (
    relid         oid
,   indexrelid    oid
,   schemaname    name
,   relname       name
,   indexrelname  name
,   idx_scan      bigint
,   idx_tup_read  bigint
,   idx_tup_fetch bigint
)
;
CREATE TABLE pg_stat_user_tables (
    relid               oid
,   schemaname          name
,   relname             name
,   seq_scan            bigint
,   seq_tup_read        bigint
,   idx_scan            bigint
,   idx_tup_fetch       bigint
,   n_tup_ins           bigint
,   n_tup_upd           bigint
,   n_tup_del           bigint
,   n_tup_hot_upd       bigint
,   n_live_tup          bigint
,   n_dead_tup          bigint
,   n_mod_since_analyze bigint
,   last_vacuum         timestamp with time zone
,   last_autovacuum     timestamp with time zone
,   last_analyze        timestamp with time zone
,   last_autoanalyze    timestamp with time zone
,   vacuum_count        bigint
,   autovacuum_count    bigint
,   analyze_count       bigint
,   autoanalyze_count   bigint
)
;
CREATE TABLE pg_stat_wal_receiver (
    pid                   integer
,   status                text
,   receive_start_lsn     pg_lsn
,   receive_start_tli     integer
,   received_lsn          pg_lsn
,   received_tli          integer
,   last_msg_send_time    timestamp with time zone
,   last_msg_receipt_time timestamp with time zone
,   latest_end_lsn        pg_lsn
,   latest_end_time       timestamp with time zone
,   slot_name             text
,   conninfo              text
)
;
CREATE TABLE pg_stat_xact_all_tables (
    relid         oid
,   schemaname    name
,   relname       name
,   seq_scan      bigint
,   seq_tup_read  bigint
,   idx_scan      bigint
,   idx_tup_fetch bigint
,   n_tup_ins     bigint
,   n_tup_upd     bigint
,   n_tup_del     bigint
,   n_tup_hot_upd bigint
)
;
CREATE TABLE pg_stat_xact_sys_tables (
    relid         oid
,   schemaname    name
,   relname       name
,   seq_scan      bigint
,   seq_tup_read  bigint
,   idx_scan      bigint
,   idx_tup_fetch bigint
,   n_tup_ins     bigint
,   n_tup_upd     bigint
,   n_tup_del     bigint
,   n_tup_hot_upd bigint
)
;
CREATE TABLE pg_stat_xact_user_functions (
    funcid     oid
,   schemaname name
,   funcname   name
,   calls      bigint
,   total_time double precision
,   self_time  double precision
)
;
CREATE TABLE pg_stat_xact_user_tables (
    relid         oid
,   schemaname    name
,   relname       name
,   seq_scan      bigint
,   seq_tup_read  bigint
,   idx_scan      bigint
,   idx_tup_fetch bigint
,   n_tup_ins     bigint
,   n_tup_upd     bigint
,   n_tup_del     bigint
,   n_tup_hot_upd bigint
)
;
CREATE TABLE pg_statio_all_indexes (
    relid         oid
,   indexrelid    oid
,   schemaname    name
,   relname       name
,   indexrelname  name
,   idx_blks_read bigint
,   idx_blks_hit  bigint
)
;
CREATE TABLE pg_statio_all_sequences (
    relid      oid
,   schemaname name
,   relname    name
,   blks_read  bigint
,   blks_hit   bigint
)
;
CREATE TABLE pg_statio_all_tables (
    relid           oid
,   schemaname      name
,   relname         name
,   heap_blks_read  bigint
,   heap_blks_hit   bigint
,   idx_blks_read   bigint
,   idx_blks_hit    bigint
,   toast_blks_read bigint
,   toast_blks_hit  bigint
,   tidx_blks_read  bigint
,   tidx_blks_hit   bigint
)
;
CREATE TABLE pg_statio_sys_indexes (
    relid         oid
,   indexrelid    oid
,   schemaname    name
,   relname       name
,   indexrelname  name
,   idx_blks_read bigint
,   idx_blks_hit  bigint
)
;
CREATE TABLE pg_statio_sys_sequences (
    relid      oid
,   schemaname name
,   relname    name
,   blks_read  bigint
,   blks_hit   bigint
)
;
CREATE TABLE pg_statio_sys_tables (
    relid           oid
,   schemaname      name
,   relname         name
,   heap_blks_read  bigint
,   heap_blks_hit   bigint
,   idx_blks_read   bigint
,   idx_blks_hit    bigint
,   toast_blks_read bigint
,   toast_blks_hit  bigint
,   tidx_blks_read  bigint
,   tidx_blks_hit   bigint
)
;
CREATE TABLE pg_statio_user_indexes (
    relid         oid
,   indexrelid    oid
,   schemaname    name
,   relname       name
,   indexrelname  name
,   idx_blks_read bigint
,   idx_blks_hit  bigint
)
;
CREATE TABLE pg_statio_user_sequences (
    relid      oid
,   schemaname name
,   relname    name
,   blks_read  bigint
,   blks_hit   bigint
)
;
CREATE TABLE pg_statio_user_tables (
    relid           oid
,   schemaname      name
,   relname         name
,   heap_blks_read  bigint
,   heap_blks_hit   bigint
,   idx_blks_read   bigint
,   idx_blks_hit    bigint
,   toast_blks_read bigint
,   toast_blks_hit  bigint
,   tidx_blks_read  bigint
,   tidx_blks_hit   bigint
)
;
CREATE TABLE pg_statistic (
    starelid    oid      NOT NULL
,   staattnum   smallint NOT NULL
,   stainherit  boolean  NOT NULL
,   stanullfrac real     NOT NULL
,   stawidth    integer  NOT NULL
,   stadistinct real     NOT NULL
,   stakind1    smallint NOT NULL
,   stakind2    smallint NOT NULL
,   stakind3    smallint NOT NULL
,   stakind4    smallint NOT NULL
,   stakind5    smallint NOT NULL
,   staop1      oid      NOT NULL
,   staop2      oid      NOT NULL
,   staop3      oid      NOT NULL
,   staop4      oid      NOT NULL
,   staop5      oid      NOT NULL
,   stanumbers1 real[]
,   stanumbers2 real[]
,   stanumbers3 real[]
,   stanumbers4 real[]
,   stanumbers5 real[]
,   stavalues1  anyarray
,   stavalues2  anyarray
,   stavalues3  anyarray
,   stavalues4  anyarray
,   stavalues5  anyarray
)
;
CREATE TABLE pg_statistic_ext (
    stxrelid        oid             NOT NULL
,   stxname         name            NOT NULL
,   stxnamespace    oid             NOT NULL
,   stxowner        oid             NOT NULL
,   stxkeys         smallint[]      NOT NULL
,   stxkind         "char"[]        NOT NULL
,   stxndistinct    pg_ndistinct
,   stxdependencies pg_dependencies
)
;
CREATE TABLE pg_stats (
    schemaname             name
,   tablename              name
,   attname                name
,   inherited              boolean
,   null_frac              real
,   avg_width              integer
,   n_distinct             real
,   most_common_vals       anyarray
,   most_common_freqs      real[]
,   histogram_bounds       anyarray
,   correlation            real
,   most_common_elems      anyarray
,   most_common_elem_freqs real[]
,   elem_count_histogram   real[]
)
;
CREATE TABLE pg_subscription (
    subdbid         oid     NOT NULL
,   subname         name    NOT NULL
,   subowner        oid     NOT NULL
,   subenabled      boolean NOT NULL
,   subconninfo     text    NOT NULL
,   subslotname     name    NOT NULL
,   subsynccommit   text    NOT NULL
,   subpublications text[]  NOT NULL
)
;
CREATE TABLE pg_subscription_rel (
    srsubid    oid    NOT NULL
,   srrelid    oid    NOT NULL
,   srsubstate "char" NOT NULL
,   srsublsn   pg_lsn NOT NULL
)
;
CREATE TABLE pg_tables (
    schemaname  name
,   tablename   name
,   tableowner  name
,   tablespace  name
,   hasindexes  boolean
,   hasrules    boolean
,   hastriggers boolean
,   rowsecurity boolean
)
;
CREATE TABLE pg_tablespace (
    spcname    name      NOT NULL
,   spcowner   oid       NOT NULL
,   spcacl     aclitem[]
,   spcoptions text[]
)
;
CREATE TABLE pg_timezone_abbrevs (
    abbrev     text
,   utc_offset interval
,   is_dst     boolean
)
;
CREATE TABLE pg_timezone_names (
    name       text
,   abbrev     text
,   utc_offset interval
,   is_dst     boolean
)
;
CREATE TABLE pg_transform (
    trftype    oid     NOT NULL
,   trflang    oid     NOT NULL
,   trffromsql regproc NOT NULL
,   trftosql   regproc NOT NULL
)
;
CREATE TABLE pg_trigger (
    tgrelid        oid          NOT NULL
,   tgname         name         NOT NULL
,   tgfoid         oid          NOT NULL
,   tgtype         smallint     NOT NULL
,   tgenabled      "char"       NOT NULL
,   tgisinternal   boolean      NOT NULL
,   tgconstrrelid  oid          NOT NULL
,   tgconstrindid  oid          NOT NULL
,   tgconstraint   oid          NOT NULL
,   tgdeferrable   boolean      NOT NULL
,   tginitdeferred boolean      NOT NULL
,   tgnargs        smallint     NOT NULL
,   tgattr         smallint[]   NOT NULL
,   tgargs         bytea        NOT NULL
,   tgqual         pg_node_tree
,   tgoldtable     name
,   tgnewtable     name
)
;
CREATE TABLE pg_ts_config (
    cfgname      name NOT NULL
,   cfgnamespace oid  NOT NULL
,   cfgowner     oid  NOT NULL
,   cfgparser    oid  NOT NULL
)
;
CREATE TABLE pg_ts_config_map (
    mapcfg       oid     NOT NULL
,   maptokentype integer NOT NULL
,   mapseqno     integer NOT NULL
,   mapdict      oid     NOT NULL
)
;
CREATE TABLE pg_ts_dict (
    dictname       name NOT NULL
,   dictnamespace  oid  NOT NULL
,   dictowner      oid  NOT NULL
,   dicttemplate   oid  NOT NULL
,   dictinitoption text
)
;
CREATE TABLE pg_ts_parser (
    prsname      name    NOT NULL
,   prsnamespace oid     NOT NULL
,   prsstart     regproc NOT NULL
,   prstoken     regproc NOT NULL
,   prsend       regproc NOT NULL
,   prsheadline  regproc NOT NULL
,   prslextype   regproc NOT NULL
)
;
CREATE TABLE pg_ts_template (
    tmplname      name    NOT NULL
,   tmplnamespace oid     NOT NULL
,   tmplinit      regproc NOT NULL
,   tmpllexize    regproc NOT NULL
)
;
CREATE TABLE pg_type (
    typname        name         NOT NULL
,   typnamespace   oid          NOT NULL
,   typowner       oid          NOT NULL
,   typlen         smallint     NOT NULL
,   typbyval       boolean      NOT NULL
,   typtype        "char"       NOT NULL
,   typcategory    "char"       NOT NULL
,   typispreferred boolean      NOT NULL
,   typisdefined   boolean      NOT NULL
,   typdelim       "char"       NOT NULL
,   typrelid       oid          NOT NULL
,   typelem        oid          NOT NULL
,   typarray       oid          NOT NULL
,   typinput       regproc      NOT NULL
,   typoutput      regproc      NOT NULL
,   typreceive     regproc      NOT NULL
,   typsend        regproc      NOT NULL
,   typmodin       regproc      NOT NULL
,   typmodout      regproc      NOT NULL
,   typanalyze     regproc      NOT NULL
,   typalign       "char"       NOT NULL
,   typstorage     "char"       NOT NULL
,   typnotnull     boolean      NOT NULL
,   typbasetype    oid          NOT NULL
,   typtypmod      integer      NOT NULL
,   typndims       integer      NOT NULL
,   typcollation   oid          NOT NULL
,   typdefaultbin  pg_node_tree
,   typdefault     text
,   typacl         aclitem[]
)
;
CREATE TABLE pg_user (
    usename      name
,   usesysid     oid
,   usecreatedb  boolean
,   usesuper     boolean
,   userepl      boolean
,   usebypassrls boolean
,   passwd       text
,   valuntil     abstime
,   useconfig    text[]
)
;
CREATE TABLE pg_user_mapping (
    umuser    oid    NOT NULL
,   umserver  oid    NOT NULL
,   umoptions text[]
)
;
CREATE TABLE pg_user_mappings (
    umid      oid
,   srvid     oid
,   srvname   name
,   umuser    oid
,   usename   name
,   umoptions text[]
)
;
CREATE TABLE pg_views (
    schemaname name
,   viewname   name
,   viewowner  name
,   definition text
)
;
