import { createSlice } from "@reduxjs/toolkit";
import Resource from "../../api/resources/Resource.js";

const reduceFn = (map, type) => {
  map[type] = {};
  return map;
};
const reduceFnInitial = {};
export const initialState = Resource.resourceTypes.reduce(
  reduceFn,
  reduceFnInitial
);

const tagSlice = createSlice({
  name: "resourceTags",
  // initialState is a map between each resource type to an empty set.
  initialState: initialState,
  reducers: {
    toggle: (state, action) => {
      const { type, tag } = action.payload;
      const set = state[type];
      if (set[tag]) {
        delete set[tag];
      } else {
        set[tag] = true;
      }
    },
  },
});

export const toggleTag = tagSlice.actions.toggle;

export default tagSlice.reducer;
