import React, { useState } from "react";
import { toast } from "react-toastify";
import { useAuthContex } from "../contex/AuthContex";
import { useMutation } from "@tanstack/react-query";

interface login {
  refElement: React.Ref<HTMLDivElement>;
}

interface requset {
  email: string;
  password: string;
}
const login = async (data: requset) => {
  const res = await fetch("/api/auth/login", {
    method: "POST",
    body: JSON.stringify(data),
  });
  return await res.json();
};
const signIn = async (data: requset) => {
  const res = await fetch("/api/auth/signup", {
    method: "POST",
    body: JSON.stringify(data),
  });
  return await res.json();
};
export const Login = (props: login) => {
  const style ="h-screen w-1/2 bg-bg absolute  top-0 right-0 bg-opacity-90 overflow-hidden   backdrop:blur";
  const mutaion = useMutation({
    mutationKey: ["login"],
    mutationFn: (reqest: requset) => login(reqest),
  });
  const SignInMutaion = useMutation({
    mutationKey: ["signup"],
    mutationFn: (reqest: requset) => signIn(reqest),
    onSuccess(data, variables, context) {
      console.log(data, variables, context);
    },
  });
  const [formData, setformData] = useState({ Email: "", Password: "" });
  const [formState, setformState] = useState(false);
  const auth = useAuthContex();
  const handleFrom = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setformData((preves) => {
      return {
        ...preves,
        [name]: value,
      };
    });
  };
  const onSubmitForm = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const requset: requset = {
      email: formData.Email,
      password: formData.Password,
    };
    if (!formState) {
      console.log(formState);
      mutaion.mutate(requset);
    } else {
      SignInMutaion.mutate(requset);
    }
    if (mutaion.isSuccess) {
      toast("login succesful", {
        position: "top-right",
        autoClose: 5000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: "dark",
      });
      auth?.setAuthstate(true);
      if (formState == false) {
        localStorage.setItem("user", JSON.stringify(mutaion.data));
      }
    }
  };
  return (
    <div id="login" className={style} ref={props.refElement}>
      <div className="flex justify-center items-center h-full overflow-hidden">
        <form className="flex flex-col h-1/2 w-1/2 " onSubmit={onSubmitForm}>
          <h2 className="text-2xl text-text-hed my-16  text-center ">
            {formState ? "Sign Up" : "Login In"}
          </h2>
          <label
            htmlFor="email"
            className="tx-sm  font-semibold leading-6 text-text-hed"
          >
            Email
          </label>
          <input
            className="text-text-hed bg-bg border-boder bg-opacity-0 backdrop:blur p-2 border-2 my-8"
            onChange={(e) => handleFrom(e)}
            type="text"
            name="Email"
            id="email"
          />
          <label
            htmlFor="password"
            className="tx-sm  font-semibold leading-6 text-text-hed"
          >
            Password
          </label>
          <input
            className="text-text-hed bg-bg border-boder bg-opacity-0 backdrop:blur p-2 border-2 my-8"
            type="password"
            name="Password"
            id="password"
          />
          <button
            id="isLogin"
            className="tx-sm  font-semibold leading-6 text-text-hed mt-8"
          >
            {formState ? "Sign Up" : "Login In"}
          </button>
          <button
            type="button"
            className="tx-sm  font-semibold leading-6 text-link mt-8"
            onClick={() => setformState(!formState)}
          >
            {!formState
              ? "Not have account? Sign Up"
              : "Already Have account! Login In"}
          </button>
        </form>
      </div>
    </div>
  );
};
