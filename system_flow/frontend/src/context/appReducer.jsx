"use client";

const subtasks = {
  id: null,
  task_id: null,
  title: null,
  created_at: null,
  deadline: null,
  completed_at: null,
};

export const initialState = [];

export const AppReducer = (state, action) => {
  let { payload } = action;
  switch (action.type) {
    case "INIT_TASK":
      return [...state, ...payload];
    case "ADD_TASK":
      return [payload, ...state];
    case "DELETE_TASK":
      return state.filter((data) => data.id != payload.id);
    case "UPDATE_TASK":
      return state.map((el) => {
        if (el.id == payload.id) {
          el["title"] = payload.title;
          el["deadline"] = payload.deadline;
          return el;
        }
        return el;
      });
    case "COMPLETE_TASK":
      return state.map((el) => {
        if (el.id == payload.id) {
          el["completed_at"] = new Date();
          return el;
        }
        return el;
      });
    case "ADD_SUBTASK":
      return state.map((el) => {
        if (el.id == payload.task_id) {
          el["subtasks"] = [...el["subtasks"], payload];
          return el;
        }
        return el;
      });
    case "UPDATE_SUBTASK":
      return state.map((el) => {
        if (el.id == payload.task_id) {
          el.subtasks.map((e) => {
            if (e.id == payload.id) {
              e.title = payload.title;
              e.deadline = payload.deadline;
              return e;
            }
            return e;
          });
          return el;
        }
        return el;
      });
    case "COMPLETE_SUBTASK":
      return state.map((el) => {
        if (el.id == payload.task_id) {
          el.subtasks.map((e) => {
            if (e.id == payload.id) {
              e.completed_at = new Date();
              return e;
            }
            return e;
          });
          return el;
        }
        return el;
      });
    case "DELETE_SUBTASK":
      return state.map((el) => {
        if (el.id == payload.task_id) {
          el.subtasks = [...el.subtasks.filter((e) => e.id != payload.id)];
          return el;
        }
        return el;
      });
    default:
      return state;
  }
};
