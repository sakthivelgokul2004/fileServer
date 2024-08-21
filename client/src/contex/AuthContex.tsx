import React, { createContext, useContext, useState } from "react"
interface AuthContextType {
  authState : boolean,
  setAuthstate :React.Dispatch<React.SetStateAction<boolean>>
}

const AuthContex = createContext<AuthContextType |null>(null)


export function useAuthContex() {
  return useContext(AuthContex)
}

export const AuthContexProvider = (props: { children: React.ReactNode }) => {
  const [authState, setauthState] = useState(Boolean)
  const obj : AuthContextType ={
    authState:authState,
    setAuthstate: setauthState
  }
  return (
    <>
      <AuthContex.Provider value={obj}>
        {props.children}
      </AuthContex.Provider>
    </>
  )
}
