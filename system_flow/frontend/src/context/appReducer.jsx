"use client";

const subtasks = {
  id: null,
  task_id: null,
  title: null,
  created_at: null,
  deadline: null,
  completed_at: null,
};

export const initialState = {
  id: "",
  title: "",
  created_at: "",
  deadline: "",
  completed_at: "",
  subtasks,
};

export const AppReducer = (state, action) => {
  switch (action.type) {
    case "USER_INPUT":
      const { payload } = action;
      return {
        ...state,
        ...payload,
      };
    case "SUBMIT_FORM":
      return { ...state, isLoading: true };
    case "SUBMIT_DONE":
      return { ...state, isLoading: false };
    default:
      return state;
  }
};
