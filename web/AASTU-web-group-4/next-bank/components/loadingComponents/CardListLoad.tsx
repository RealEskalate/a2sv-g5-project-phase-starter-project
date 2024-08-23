import { CardListSkeleton } from "./CardListSkeleton"

const CardListLoad = () => {
  return (
    <div className="cardlist lg:w-[730px] md:w-[487px] sm-w-[325] my-6">
        <CardListSkeleton />
        <CardListSkeleton />
        <CardListSkeleton />
        <CardListSkeleton />
    </div>
  )
}

export default CardListLoad
