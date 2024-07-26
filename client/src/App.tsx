
import React, { useRef } from "react"
function App() {
  const refs = useRef<HTMLFormElement>(null);

  const inputref = useRef<HTMLInputElement>(null);
  function news(e: React.FormEvent<HTMLFormElement>) {
    console.log("NOTU")
    if (inputref.current?.files != null) {
      const file =inputref.current?.files[0]
      console.log(file)
      const form = new FormData()
      form.append("image",file)
      fetch("http://localhost:8080/upload", {
        method: "POST",
        body: form
      }
      )
    }
    console.log("HI")
    e.preventDefault()
  }
//  {//encType="multipart/form-data" method="POST" action="http://localhost:8080/upload" id="image" ref={refs} >}
  return (
    <>
    <h5>hi</h5>
      <form ref={refs} onSubmit={news}>
        <input type="file" id="image" name="image"ref={inputref} />
        <button type="submit" onClick={() => (console.log("hi"))}>summit</button>
      </form>
    </>
  )
}

export default App
