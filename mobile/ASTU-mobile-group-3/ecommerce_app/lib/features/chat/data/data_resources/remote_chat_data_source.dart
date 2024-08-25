
import 'package:http/http.dart' as http;


abstract class RemoteChatDataSource {



}




class RemoteChatDataSourceImp implements RemoteChatDataSource {
  final http.Client client;


  RemoteChatDataSourceImp(this.client);

  
}