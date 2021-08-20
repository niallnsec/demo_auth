package auth

import (
	_ "github.com/BurntSushi/toml"
	_ "github.com/aws/aws-sdk-go"
	_ "github.com/fatih/color"
	_ "github.com/gofrs/uuid"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/gorilla/csrf"
	_ "github.com/gorilla/handlers"
	_ "github.com/hashicorp/vault/api"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/julienschmidt/httprouter"
	_ "github.com/justinas/alice"
	_ "github.com/mailru/easyjson"
	_ "github.com/microcosm-cc/bluemonday"
	_ "github.com/mssola/user_agent"
	_ "github.com/ogier/pflag"
	_ "github.com/otiai10/copy"
	_ "github.com/prometheus/client_golang/prometheus"
	_ "github.com/rs/cors"
	_ "github.com/rs/zerolog"
	_ "github.com/unrolled/secure"
	_ "github.com/valyala/fasthttp"
)

func Run() {

}
