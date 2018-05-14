package channels

//Note: Only the sender should close a channel,
//never the receiver. Sending on a closed channel
//will cause a panic.
//
//Another note: Channels aren’t like files
//you don’t usually need to close them. Closing is only
//necessary when the receiver must be told there are no
//more values coming, such as to terminate a range loop.

func Channel(){

}