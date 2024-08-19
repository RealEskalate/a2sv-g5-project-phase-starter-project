import * as React from "react"
const CorporateLoans = (props) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={45}
    height={45}
    fill="none"
    {...props}
  >
    <circle cx={22.5} cy={22.5} r={22.5} fill="#FFF5D9" />
    <g fill="#FFBB38" clipPath="url(#a)">
      <path d="M30.5 17.75h-3.75V17c0-.827-.673-1.5-1.5-1.5h-4.5c-.827 0-1.5.673-1.5 1.5v.75H15.5c-.827 0-1.5.673-1.5 1.5v2.25c0 .827.673 1.5 1.5 1.5h6v-.375c0-.207.168-.375.375-.375h2.25c.207 0 .375.168.375.375V23h6c.827 0 1.5-.673 1.5-1.5v-2.25c0-.827-.673-1.5-1.5-1.5Zm-5.25 0h-4.5V17h4.5v.75ZM31.792 23.405a.375.375 0 0 0-.393.037 1.48 1.48 0 0 1-.899.308h-6v1.125a.375.375 0 0 1-.375.375h-2.25a.375.375 0 0 1-.375-.375V23.75h-6a1.48 1.48 0 0 1-.899-.308.375.375 0 0 0-.601.299V29c0 .827.673 1.5 1.5 1.5h15c.827 0 1.5-.673 1.5-1.5v-5.26a.375.375 0 0 0-.208-.335Z" />
    </g>
    <defs>
      <clipPath id="a">
        <path fill="#fff" d="M14 14h18v18H14z" />
      </clipPath>
    </defs>
  </svg>
)
export default CorporateLoans
