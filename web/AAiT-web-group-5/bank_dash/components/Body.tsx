import CardS from "./CardS";
import Top from "./Top";
import Topleft from "./Topleft";

export default function Body() {
  return (
    <div
      className=" bg-slate-50 flex-col overflow-auto"
      style={{
        position: "fixed",
        top: 0,
        left: 0,
        width: "100%",
        height: "100vh",
      }}
    >
      <Top />
      <div className="flex justify-center">
        <CardS />
      </div>
    </div>
  );
}
