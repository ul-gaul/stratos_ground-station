

const units = {
  get nanosecond() { return this.microsecond / 1000 },
  get microsecond() { return  this.millisecond / 1000 },
  millisecond: 1,
  get second() { return this.millisecond * 1000 },
  get minute() { return this.second * 60 },
  get hour() { return this.minute * 60 },
  get day() { return this.hour * 24 },
  get month() { return this.year / 12 },
  get year() { return this.day * 365.25 },
};

export default units;

