import React from "react";
import Center from "./Center";
import Bottom from "./Bottom";

const Container = () => {
  return (
    <section className="w-[96%] flex flex-col grow gap-6 p-8 pt-6 xs:pt-20 xs:px-4 sm:px-8 sm:pt-6">
      <Center />
      <Bottom />
    </section>
  );
};

export default Container;
