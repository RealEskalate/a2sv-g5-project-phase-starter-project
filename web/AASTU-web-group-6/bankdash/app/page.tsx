import React from "react";
import Container from "./components/Dashboard/Container";
import { getServerSession } from "next-auth";
import { options } from "./api/auth/[...nextauth]/options";
import LoginForm from "./components/Forms/LoginForm";
const Home = () => {
  // const session = await getServerSession(options);
  return (
    <Container />
    // <>
    //   {/* {session && <Container />} */}
    //   {/* {!session && <LoginForm />} */}
    // </>
  );
};
export default Home;
