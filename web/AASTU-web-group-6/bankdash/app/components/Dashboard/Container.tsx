import React from "react";
import Center from "./Center";
import Right from "./Right";
import Bottom from "./Bottom";

const Container = () => {
  return (
    <section className="w-full flex flex-col grow gap-6 p-8">
      <div className="w-full flex gap-8">
        <Center />
        <Right />
      </div>
      <Bottom />
    </section>
  );
};

export default Container;
