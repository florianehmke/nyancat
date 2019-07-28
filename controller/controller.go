package controller

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	"github.com/florianehmke/nyancat/api"
)

type Controller struct {
	run bool
	mr  *MinerRegistry
	wg  sync.WaitGroup
}

func NewController() *Controller {
	return &Controller{mr: NewMinerRegistry(), run: true}
}

func (c *Controller) Start() {
	c.wg.Add(2)
	go c.RequestCats()
	go c.ServeManagementApi()

	c.wg.Wait()
}

func (c *Controller) RequestCats() {
	for c.run == true {
		miners := c.mr.Miners()
		for _, m := range miners {
			requestCatFromServer(m)
		}

		time.Sleep(time.Second * 1)
	}
	c.wg.Done()
}

func requestCatFromServer(s string) {
	conn, err := grpc.Dial(s, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewMinerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.MineCat(ctx, &api.MineRequest{Id: uuid.New().String()})
	if err != nil {
		log.Fatalf("could not mine cat: %v", err)
	}
	log.Printf("%s: %s", r.Id, r.Cat)
}
