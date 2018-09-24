import http from "k6/http";
import { sleep } from "k6";

export default function() {
    var params =  { headers: { "Content-Type": "application/json" } }
    http.post("http://localhost:8080/blog/getAllArticle", "", params);
};