import React from "react";
import Center from "./Center";
import Bottom from "./Bottom";

const Container = () => {
  return (
    <section className="w-[96%] flex flex-col grow gap-6 p-8  xxs:px-4 sm:px-8 md:pt-6 ">
      <Center />
      <Bottom />
    </section>
  );
};

export default Container;
