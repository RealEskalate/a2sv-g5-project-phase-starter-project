import Topleft from "./Topleft";
import Topright from "./Topright";

export default function Top() {
  return (
    <div className="flex justify-between w-full mb-5">
      <Topright topic="OverView" />
      <Topleft />
    </div>
  );
}
