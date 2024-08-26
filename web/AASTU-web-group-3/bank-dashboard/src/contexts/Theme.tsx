import { createContext,useContext,useState,useEffect ,ReactNode, Children} from "react";


interface themeType{
    theme: "dark" | "light",
    toggleTheme: () => void,
    setTheme: (theme: "dark" | "light") => void;
}
const ThemeContext  = createContext<themeType | undefined>(undefined)

export const ThemeProvider = ({children}:{children:ReactNode}) =>{
    const [theme,setTheme] = useState<'light' | 'dark'>('light')


    useEffect(() => {
        const savedTheme = localStorage.getItem('theme') as "dark" | "light"
        if(savedTheme){
            setTheme(savedTheme)
            document.documentElement.classList.add(savedTheme)
        }

    },[])

    
    const toggleTheme = () => {
        const newTheme = theme === 'light'?'dark':'light'
        setTheme(newTheme)
        document.documentElement.classList.remove(theme)
        document.documentElement.classList.add(newTheme)
        localStorage.setItem('theme',newTheme)
    } 

    return (
        <ThemeContext.Provider value={{theme, setTheme,toggleTheme}}>
            {children}
        </ThemeContext.Provider>
    )
    
}

export const useTheme = () => {
    const context = useContext(ThemeContext);
    if (!context){
        throw new Error('useTheme must be used within a ThemeProvider')
    }
    return context
}