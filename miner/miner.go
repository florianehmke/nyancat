package miner

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/florianehmke/nyancat/api"
)

var nyans = []string{
	"[^._.^]ﾉ彡",
	"~=[,,_,,]:3",
	"(ﾐㆁ ﻌ ㆁﾐ)∫",
	"=^._.^= ∫ ",
	"ฅ(＾・ω・＾ฅ)",
	"（＾・ω・＾✿） ",
	"（＾・ω・＾❁） ",
	"(=^･ω･^=)",
	"(^・x・^)",
	"(=^･ｪ･^=))ﾉ彡☆ ",
	"(^._.^)ﾉ☆",
	"( _ _).oO ",
	"(=‐ω‐=)",
	"(=｀ω´=)",
	"(=｀ェ´=)",
	"（=´∇｀=） ",
	"(=´∇｀=) ",
	"(=^ ◡ ^=) ",
	"(=^-ω-^=) ",
	"(＾º◡º＾❁) ",
	"(^≗ω≗^)ฅ",
	"(•ㅅ•❀) ฅ",
	"(＊☉౪ ⊙｡)",
	"(ﾐㆁ ﻌ ㆁﾐ)∫ ",
	"[^._.^]ﾉ彡  ",
	"~=[,,_,,]:3 ",
	"~=[,,_,,]^‥^ ",
	"^._.^= ∫ ",
	"(=ච ω ච=)",
	"( ΦωΦ )",
	"ฅ^•ﻌ•^ ",
	"=＾ȏ⋏ȏ＾=  ",
	"(=ච㉨ච=)",
	"=( ^>ܫ< ^)= ",
	"~(´^ゝω･)",
	"＾・ﻌ・＾ ",
	"ฅ ( ◕ ﻌ ◕ )",
	"(ﾐoᆽoﾐ)",
	"(=^・ ◡・ ^=)",
	"(=ඒᆽඒ=)",
	"(ﾐΦ ﻌ Φﾐ)∫ ",
	"ﾐ|ዎ ﻌ ዎ|ﾐ  ",
	"(₌oᆽo₌)",
	"(=•ω•=)",
	"Ɛ:[,,_,,]=~",
	"(＾*･.･*＾)",
	"(=･ｪ･=)",
	"*:･ﾟ✧(=✪ ᆺ ✪=)*:･ﾟ✧ ",
	"=＾ᵒ⋏ᵒ＾=  ",
	"(= ච㉨ච=)",
	"(ﾐටᆽටﾐ)",
	"(=◕ᆽ◕= ✿)",
}

func MineCat() string {
	rand.Seed(time.Now().Unix())
	message := nyans[rand.Intn(len(nyans))]
	return message
}

const (
	port = ":50051"
)

type server struct{}

func (s *server) MineCat(ctx context.Context, in *api.MineRequest) (*api.MineReply, error) {
	log.Printf("Processing Request: %v", in.Id)
	return &api.MineReply{Id: in.Id, Cat: MineCat()}, nil
}

func Serve() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterMinerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
