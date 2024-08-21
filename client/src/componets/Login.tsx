
import React, { useState } from 'react'

interface login {
  refElement: React.Ref<HTMLDivElement>

}

export const Login = (props: login) => {
  const style = "  h-screen w-1/2 bg-bg absolute  top-0 right-0 bg-opacity-90 z-0  backdrop:blur"
  const [formData, setformData] = useState({ Email: "", Password: "" })
  const [formState, setformState] = useState(false)
  const handleFrom = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setformData((preves) => {
      return {
        ...preves,
        [name]: value
      }
    })
  }
  const onSubmitForm = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const url = formState ? "/auth/signup" : "/auth/login"
    console.log(formData)
    const requset = {
      email: formData.Email,
      password: formData.Password,
    }
    try {
      fetch(url, {
        method: "POST",
        body: JSON.stringify(requset)
      }).then((res) => console.log(res))
        .catch((err) => console.log(err))
    } catch (error) {
      console.log(error)
    }


  }
  return (
    <div id="login" className={style} ref={props.refElement}>
      <div className="flex justify-center items-center h-full " >
        <form className="flex flex-col h-1/2 w-1/2 " onSubmit={onSubmitForm}>
          <h2 className="text-2xl text-text-hed my-16  text-center ">{formState ? "Sign Up" : "Login In"}</h2>
          <label htmlFor="email" className="tx-sm  font-semibold leading-6 text-text-hed">Email</label>
          <input className="text-text-hed bg-bg border-boder bg-opacity-0 backdrop:blur p-2 border-2 my-8" onChange={(e) => handleFrom(e)} type="text" name="Email" id="email" />
          <label htmlFor="password" className="tx-sm  font-semibold leading-6 text-text-hed">Password</label>
          <input className="text-text-hed bg-bg border-boder bg-opacity-0 backdrop:blur p-2 border-2 my-8" type="password" name="Password" id="password" />
          <button className="tx-sm  font-semibold leading-6 text-text-hed mt-8" >{formState ? "Sign Up" : "Login In"}</button>
          <button className="tx-sm  font-semibold leading-6 text-text-hed mt-8" onClick={() => setformState(!formState)}>{formState ? "Login In" : "Sign Up"}</button>
        </form>
      </div>
    </div>
  )
}
