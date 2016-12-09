package main

import (
    "fmt"
    "time"
    "strconv"
)

func makeCakeAndSend(cs chan string, flavor string, count int) {
    for i := 1; i <= count; i++ {
        cakename := flavor + " Cake " + strconv.Itoa(i)
        cs <- cakename //send a strawberry cake
        time.Sleep(200*time.Millisecond);
   
    }   
    close(cs)
}

func receiveCakeAndPack(strbry_cs chan string, choco_cs chan string) {
    strbry_closed, choco_closed := false, false

    for {
        //if both channels are closed then we can stop
        if (strbry_closed && choco_closed) { return }
        fmt.Println("Waiting for a new cake ...")
        select {
        case cakename,strbry_ok := <-strbry_cs:
            if (!strbry_ok) {
                strbry_closed = true
                fmt.Println(" ... Strawberry channel closed!")
            } else {
                fmt.Println("Received from Strawberry channel.  Now packing", cakename)
            }   
        case cakename, choco_ok := <-choco_cs:
            if (!choco_ok) {
                choco_closed = true
                fmt.Println(" ... Chocolate channel closed!")
            } else {
                fmt.Println("Received from Chocolate channel.  Now packing", cakename)
            }   
        }   
    }   
}

func main() {
    strbry_cs := make(chan string)
    choco_cs := make(chan string)

     //time.Sleep(400 * time.Millisecond);
    //two cake makers
    go makeCakeAndSend(choco_cs, "Chocolate", 3)  //make 3 chocolate cakes and send
        time.Sleep(200*1000);
    go makeCakeAndSend(strbry_cs, "Strawberry", 3)  //make 3 strawberry cakes and send
        time.Sleep(400*1000);

    //one cake receiver and packer
    go receiveCakeAndPack(strbry_cs, choco_cs)  //pack all cakes received on these cake channels

    //sleep for a while so that the program doesn’t exit immediately
    time.Sleep(2 * 1e9)
}
