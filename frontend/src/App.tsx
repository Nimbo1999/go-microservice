import SectionContent from "./components/section-content"

function App() {
  return (
    <main className="container">
      <header className="header">
        <h1>Test microservices</h1>
        <hr className="container__row-line" />
      </header>

      <section className="container__content">
        <SectionContent title="Sent">
          Nothing sent yet...
        </SectionContent>

        <SectionContent title="Received">
          Nothing received yet...
        </SectionContent>
      </section>

      <hr className="container__row-line" />

      <span className="container__copyright">Copyright &copy; Matheus Lopes</span>
    </main>
  )
}

export default App
