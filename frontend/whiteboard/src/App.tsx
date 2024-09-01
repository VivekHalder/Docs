import { useEffect, useState } from "react"
import ToolBar from "./components/ToolBar";
import TextArea from "./components/TextArea";

function App() {
  const [content, setContent] = useState<string>("");
  const [socket, setSocket] = useState<WebSocket | null>(null);
  const documentID = 'doc1';

  useEffect(() => {
    const ws = new WebSocket('ws://localhost:8080/ws');

    ws.onopen = () => {
      console.log("Connected to WebSocket server");
    }

    ws.onmessage = (event) => {
      const updatedDoc = JSON.parse(event.data);

      if (updatedDoc.id === documentID) {
        setContent(updatedDoc.content)
      }
    }

    setSocket(ws);

    return () => {
      ws.close();
    }

  }, []);

  const handleChange = (event: React.FormEvent<HTMLDivElement>) => {
    const value = (event.target as HTMLDivElement).innerHTML;
    setContent(value);
    if (socket) {
      const message = {
        data: {
          id: documentID,
          content: value
        }
      }
      socket.send(JSON.stringify(message));
    }
  }

  return (
    <div>
      <ToolBar />
      <TextArea value={content} onChange={handleChange} />
    </div>
  )
}

export default App
