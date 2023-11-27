import { useState } from "react";
import SectionContent from "./components/section-content"

function App() {
  const [payloads, setPayloads] = useState<any[]>([]);
  const [responses, setResponses] = useState<any[]>([]);

  const format = (data: any) => JSON.stringify(data, undefined, 4);

  return (
    <main className="container">
      <header className="header">
        <h1>Test microservices</h1>
        <hr className="container__row-line" />
        <button onClick={async () => {
          setPayloads((prev) => [...prev, null])
          try {
            const request = await fetch('http://localhost:8080', { method: 'POST' });
            const response = await request.json()
            setResponses((prev) => [...prev, response])
          } catch(err) {
            console.error(err)
          }
        }}>
          Test
        </button>
      </header>

      <section className="container__content">
        <SectionContent title="Sent">
          <div className="container__content-wrapper">
              {payloads.length < 1 ? "Nothing sent yet..." : payloads.map(payload => (
              <pre className="container__content-item">{format(payload)}</pre>
              ))}
          </div>
        </SectionContent>

        <SectionContent title="Received">
          <div className="container__content-wrapper">
            {responses.length < 1 ? "Nothing sent yet..." : responses.map(response => (
              <pre className="container__content-item">{format(response)}</pre>
              ))}
          </div>
        </SectionContent>
      </section>

      <hr className="container__row-line" />

      <span className="container__copyright">Copyright &copy; Matheus Lopes</span>
    </main>
  )
}

export default App
