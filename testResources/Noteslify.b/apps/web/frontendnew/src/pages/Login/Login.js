import { useContext } from "react";
import { Helmet } from "react-helmet";
import { ThemeContext } from "../../contexts/ThemeModeProvider";
import LoginContainer from "../../components/Container/LoginContainer";
import StickyFooter from "../../components/Footer/StickyFooter";

const Login = () => {
  const { theme } = useContext(ThemeContext);

  const dark = theme === "dark";

  return (
    <>
      <Helmet>
        <title>Noteslify | Home</title>
        <meta
          name="description"
          content="Your Privacy Friendly, Open Source. Alternative to EverNote."
        />
      </Helmet>
      {/* main */}
      <main>
      <LoginContainer />
      </main>
      {/* footer */}
      <StickyFooter dark={dark} />
    </>
  );
};

export default Login;
