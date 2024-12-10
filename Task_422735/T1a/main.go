	name := r.URL.Query().Get("name")
	if name == "" {
		conn.WriteMessage(websocket.TextMessage, []byte("Please provide a username."))
		return
	}

	clients[conn] = name
	log.Printf("Connected: %s\n", name)

	defer func() {
		delete(clients, conn)
		log.Printf("Disconnected: %s\n", name)
	}()

	for {
		select {
		case message := <-broadcast:
			err := conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("Writing message error:", err)
				return
			}
		case err := <-conn.Err():
			log.Println("Connection error:", err)
			return
		}
	}