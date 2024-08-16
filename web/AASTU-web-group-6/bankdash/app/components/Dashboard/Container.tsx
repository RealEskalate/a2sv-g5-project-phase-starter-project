import React from "react";
import Center from "./Center";
import Bottom from "./Bottom";

const Container = () => {
  return (
    <section className="w-full flex flex-col grow gap-6 p-8">
      <Center />
      <Bottom />
    </section>
  );
};

export default Container;
