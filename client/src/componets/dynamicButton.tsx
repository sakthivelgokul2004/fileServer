import { useAuthContex } from "../contex/AuthContex";

interface loginProps {
  logout: React.MouseEventHandler<HTMLButtonElement>;
  login: React.MouseEventHandler<HTMLButtonElement>;
}
const DynamicButton: React.FC<loginProps> = ({ login, logout }) => {
  const auth = useAuthContex();
  if (auth?.authState == true) {
    return (
      <>
        <button className="ml-auto text-text-hed" onClick={logout}>
          {" "}
          Logout
        </button>
      </>
    );
  }
  return (
    <>
      <button className="ml-auto text-text-hed" onClick={login}>
        {" "}
        LogIn
      </button>
    </>
  );
};

export default DynamicButton;
