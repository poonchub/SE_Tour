import { createContext, ReactNode, useContext, useState } from 'react';

interface DateContextProps {
    dateSelectedFormat: any,
    setDateSelectedFormat: (param: string) => void,
}

const DateContext = createContext<DateContextProps>({
    dateSelectedFormat: "",
    setDateSelectedFormat: () => {},
});

export const DateProvider = ({ children }: { children: ReactNode }) => {
    const [dateSelectedFormat, setDateSelectedFormat] = useState<string>("");
    return (
        <DateContext.Provider value={{ dateSelectedFormat, setDateSelectedFormat }}>
            {children}
        </DateContext.Provider>
    );
};

export const useDateContext = (): DateContextProps => {
    const context = useContext(DateContext);
    if (!context) {
        throw new Error("useDateContext must be used within a DateProvider");
    }
    return context;
};