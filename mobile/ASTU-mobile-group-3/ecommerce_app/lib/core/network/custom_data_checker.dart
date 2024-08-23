import 'package:http/http.dart' as http;
class DataConnectionChecker{
  /// Check if there is a connection by sending request to server that are guaranted that they are always online
  /// 
  /// return true if there is connection otherwise false
  final http.Client client;

  DataConnectionChecker(this.client);


  Future<bool> hasConnection(String url) async {
    bool answer = false;
    try{
      final result = await client.get(Uri.parse(url));
      if (result.statusCode == 200) {
        answer = true;
      }
    } on Exception{
      answer = false;
    }finally{
      client.close();
    }


    return answer;
  }
}