import { AuthContexProvider } from './contex/AuthContex';
import { Home } from './page/home';


function App() {
  return (
    <div className="h-screen w-screen  bg-bg">
      <div className="w-[108] flex justify-end flex-none overflow-hidden">
        <div className="z-0  top-0 flex    absolute">
          <picture className="">
            <source srcSet="./d.png"></source>
            <img src="/d.png" className="   w-[90rem] max-w-none flex-none   " />
          </picture>
        </div>
      </div>
      <AuthContexProvider>
      <Home />
      </AuthContexProvider>
    </div>


  )
}

export default App                
