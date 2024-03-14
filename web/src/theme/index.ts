const common = {
  Message: {
    borderRadius: "0.5rem",
  },
  Dialog: {
    borderRadius: "0.75rem",
  },
  Carousel: {
    dotColor: "rgba(24,160, 88,0.2)",
    dotColorActive: "#18A058",
  },
};

export const themeProvider = (other: any) => {
  return {
    ...common,
    ...other,
  };
};
