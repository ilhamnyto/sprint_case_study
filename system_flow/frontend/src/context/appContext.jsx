"use client";
import { useReducer, useMemo, createContext, useContext } from "react";

import { AppReducer, initialState } from "./appReducer";

export const AppContext = createContext();

export const useAppContext = () => useContext(AppContext);

export function AppWrapper({ children }) {
  const [state, dispatch] = useReducer(AppReducer, initialState);

  const contextValue = useMemo(() => {
    return { state, dispatch };
  }, [state, dispatch]);

  return (
    <AppContext.Provider value={contextValue}>{children}</AppContext.Provider>
  );
}
