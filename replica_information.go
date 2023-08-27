package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type replica struct {
	IO_state                 string
	SourceHost               string
	SourceUser               string
	SourcePort               int
	ConnectRetry             int
	SourceLogFile            string
	ReadSourceLogPos         int
	RelayLogfile             string
	RelayLogPos              int
	RelaySourceLogFile       string
	Replica_IO_State         string
	Replica_SQL_State        string
	ReplicateDoDb            string
	ReplicateIgnoreDB        string
	ReplicateDoTable         string
	ReplicateIgnoreTable     string
	ReplicateWildDoTable     string
	ReplicateWildIgnoreTable string
	LastErrorNum             int
	LastError                string
	SkipCounter              int
	ExecSourceLogPos         int
	RelayLogSpace            int
	UntilCond                string
	UntilLogFile             string
	UntilLogPos              int
	SourceSSLAllow           string
	SourceSSLCAFile          string
	SourceSSLCAPath          string
	SourceSSLCert            string
	SourceSSLCipher          string
	SourceSSLKey             string
	SecondsBehindSource      sql.NullString
	SourceSSLVerifyServer    string
	LastIOErrorNum           int
	LastIOError              string
	LastSQLErrorNum          int
	LastSQLError             string
	ReplicateIgnoreServerIDs string
	SourceServerId           int
	SourceUUID               string
	SourceInfoFile           string
	SQLDelay                 int
	SQLRemainingDelay        sql.NullString
	ReplicaSQLRunningState   string
	SourceRetryCount         int
	SourceBind               string
	LastIOErrorTimestamp     string
	LastSQLErrorTimestamp    string
	SourceSSLCrl             string
	SourceSSLCrlPath         string
	RetrievedGTIDSet         string
	ExecutedGTIDSet          string
	AutoPos                  int
	ReplicateRewriteDB       string
	ChannelName              string
	SourceTLSVersion         string
	SourcePublicKeyPath      string
	GetSourcePublicKey       int
	NetworkNamespace         string
}

func main() {

	db, err := sql.Open("mysql", "dbadmin:gw2345@tcp(127.0.0.1:22335)/totin_sandbox_db1")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("show replica status")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("The query results:  ", results)

	for results.Next() {
		var repl_res replica
		err = results.Scan(
			&repl_res.IO_state,
			&repl_res.SourceHost,
			&repl_res.SourceUser,
			&repl_res.SourcePort,
			&repl_res.ConnectRetry,
			&repl_res.SourceLogFile,
			&repl_res.ReadSourceLogPos,
			&repl_res.RelayLogfile,
			&repl_res.RelayLogPos,
			&repl_res.RelaySourceLogFile,
			&repl_res.Replica_IO_State,
			&repl_res.Replica_SQL_State,
			&repl_res.ReplicateDoDb,
			&repl_res.ReplicateIgnoreDB,
			&repl_res.ReplicateDoTable,
			&repl_res.ReplicateIgnoreTable,
			&repl_res.ReplicateWildDoTable,
			&repl_res.ReplicateWildIgnoreTable,
			&repl_res.LastErrorNum,
			&repl_res.LastError,
			&repl_res.SkipCounter,
			&repl_res.ExecSourceLogPos,
			&repl_res.RelayLogSpace,
			&repl_res.UntilCond,
			&repl_res.UntilLogFile,
			&repl_res.UntilLogPos,
			&repl_res.SourceSSLAllow,
			&repl_res.SourceSSLCAFile,
			&repl_res.SourceSSLCAPath,
			&repl_res.SourceSSLCert,
			&repl_res.SourceSSLCipher,
			&repl_res.SourceSSLKey,
			&repl_res.SecondsBehindSource,
			&repl_res.SourceSSLVerifyServer,
			&repl_res.LastIOErrorNum,
			&repl_res.LastIOError,
			&repl_res.LastSQLErrorNum,
			&repl_res.LastSQLError,
			&repl_res.ReplicateIgnoreServerIDs,
			&repl_res.SourceServerId,
			&repl_res.SourceUUID,
			&repl_res.SourceInfoFile,
			&repl_res.SQLDelay,
			&repl_res.SQLRemainingDelay,
			&repl_res.ReplicaSQLRunningState,
			&repl_res.SourceRetryCount,
			&repl_res.SourceBind,
			&repl_res.LastIOErrorTimestamp,
			&repl_res.LastSQLErrorTimestamp,
			&repl_res.SourceSSLCrl,
			&repl_res.SourceSSLCrlPath,
			&repl_res.RetrievedGTIDSet,
			&repl_res.ExecutedGTIDSet,
			&repl_res.AutoPos,
			&repl_res.ReplicateRewriteDB,
			&repl_res.ChannelName,
			&repl_res.SourceTLSVersion,
			&repl_res.SourcePublicKeyPath,
			&repl_res.GetSourcePublicKey,
			&repl_res.NetworkNamespace,
		)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println("Current Primary Host: ", repl_res.SourceHost)
		fmt.Println("Primary UUID:  ", repl_res.SourceUUID)
		fmt.Println("Replica IO Running:  ", repl_res.Replica_IO_State)
		fmt.Println("Replica SQL Running:  ", repl_res.Replica_SQL_State)
		fmt.Println("Retrieved GTID:  ", repl_res.RetrievedGTIDSet)
		fmt.Println("Executed GTID:  ", repl_res.ExecutedGTIDSet)
	}

}
