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
            const request = await fetch(import.meta.env.VITE_BROKER_HOST, { method: 'POST' });
            const response = await request.json()
            setResponses((prev) => [...prev, response])
          } catch(err) {
            console.error(err)
          }
        }}>
          Test Broker
        </button>

        <button onClick={async () => {
          const payload = {
            action: 'auth',
            auth: {
              email: 'admin@example.com',
              password: 'verysecret'
            }
          }

          setPayloads((prev) => [...prev, payload])
          try {
            const request = await fetch(`${import.meta.env.VITE_BROKER_HOST}/handle`, {
              method: 'POST',
              body: JSON.stringify(payload),
              headers: {
                'Content-Type': 'application/json'
              }
            });
            const response = await request.json()
            setResponses((prev) => [...prev, response])
          } catch(err) {
            console.error(err)
          }
        }}>
          Test Auth
        </button>

        <button onClick={async () => {
          const payload = {
            action: 'log',
            log: {
              name: 'user',
	            data: JSON.stringify({name: 'Matheus', email: 'matlopes1999@gmail.com'}),
            }
          }

          setPayloads((prev) => [...prev, payload])
          try {
            const request = await fetch(`${import.meta.env.VITE_BROKER_HOST}/handle`, {
              method: 'POST',
              body: JSON.stringify(payload),
              headers: {
                'Content-Type': 'application/json'
              }
            });
            const response = await request.json()
            setResponses((prev) => [...prev, response])
          } catch(err) {
            console.error(err)
          }
        }}>
          Test Logger
        </button>

        <button onClick={async () => {
          const payload = {
            name: 'user',
            data: JSON.stringify({name: 'Matheus gRPC', email: 'matlopes1999@gmail.com'}),
          }

          setPayloads((prev) => [...prev, payload])
          try {
            const request = await fetch(`${import.meta.env.VITE_BROKER_HOST}/log-grpc`, {
              method: 'POST',
              body: JSON.stringify(payload),
              headers: {
                'Content-Type': 'application/json'
              }
            });
            const response = await request.json()
            setResponses((prev) => [...prev, response])
          } catch(err) {
            console.error(err)
          }
        }}>
          Test Logger Via gRPC
        </button>

        <button onClick={async () => {
          const payload = {
            action: 'mail',
            mail: {
              from: 'me@example.com',
              to: 'you@there.com',
              subject: 'Test mail',
              message: 'Hello world!'
            }
          }

          setPayloads((prev) => [...prev, payload])
          try {
            const request = await fetch(`${import.meta.env.VITE_BROKER_HOST}/handle`, {
              method: 'POST',
              body: JSON.stringify(payload),
              headers: {
                'Content-Type': 'application/json'
              }
            });
            const response = await request.json()
            setResponses((prev) => [...prev, response])
          } catch(err) {
            console.error(err)
          }
        }}>
          Test Mail
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
