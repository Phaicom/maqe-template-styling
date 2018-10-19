new Vue({
  el: "#app",
  delimiters: ["${", "}"],
  data() {
    return {
      forums: [],
      errors: [],
      currentPage: 1,
      pageSize: 8,
      total: 0,
      loading: false,
    };
  },
  created() {
    this.loading = true;
    axios
      .get("/api/forum")
      .then(response => {
        total = response.data.length;
        min = this.pageSize * (this.currentPage - 1);
        max = this.pageSize * this.currentPage;
        this.forums = response.data.slice(min, max);
        this.loading = false;
      })
      .catch(e => {
        this.errors.push(e);
        this.loading = false;
      });
  },
  methods: {
    fromDate: date => {
      return moment(date).fromNow();
    },
    handleCurrentChange(val) {
      this.loading = true;
      axios
        .get("/api/forum")
        .then(response => {
          total = response.data.length;
          min = this.pageSize * (this.currentPage - 1);
          max = this.pageSize * this.currentPage;
          this.forums = response.data.slice(min, max);
          this.loading = false;
        })
        .catch(e => {
          this.errors.push(e);
          this.loading = false;
        });
    },
  },
});
