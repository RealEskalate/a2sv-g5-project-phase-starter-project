import Topleft from "./Topleft";
import Topright from "./Topright";

export default function Top({ topicName }: { topicName: string }) {
  return (
    <>
    
    <div className="flex flex-row justify-between w-full pb-7 mb-5 mr-5 ">
      <Topleft topic={topicName} />
      <Topright />
    </div>
    </>
  );
}
