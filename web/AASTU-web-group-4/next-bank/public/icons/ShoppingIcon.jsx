import * as React from "react"
const ShoppingIcon = (props) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={50}
    height={50}
    fill="none"
    {...props}
  >
    <circle cx={25} cy={25} r={25} fill="#FFF5D9" />
    <path
      fill="#FFBB38"
      d="m32.313 32.275-1.067-11.706a.624.624 0 0 0-.622-.569h-1.875v-1.25c0-1.005-.39-1.948-1.096-2.654a3.755 3.755 0 0 0-6.404 2.654V20h-1.875a.624.624 0 0 0-.623.569l-1.064 11.705a2.504 2.504 0 0 0 .644 1.912 2.504 2.504 0 0 0 1.846.814h9.644c.7 0 1.373-.296 1.845-.813a2.512 2.512 0 0 0 .646-1.912ZM27.498 20h-5v-1.25a2.502 2.502 0 0 1 4.27-1.771c.471.471.73 1.1.73 1.771V20Z"
    />
  </svg>
)
export default ShoppingIcon
