
import 'package:socket_io_client/socket_io_client.dart' as IO;

import '../../../../../core/constants/constants.dart';

class SocketService {
  SocketService();
  
  IO.Socket? socket  ;
  
  void connectAndListen(){
    socket ??= IO.io(Uri.parse(Urls.baseChat));

    socket!.onConnect((data)=>{
      print('connected to server')
    });
  }

  void sendMessage(String chatId,String message,String type){
        Map<String,dynamic>messageFormat = {
            'chatId': chatId,
            'content':message,
            'type': type
        };
        socket!.emit('messagej',messageFormat);
  }

  void dispose(){
    if(socket!.connected){
      socket!.disconnect();
    }
  }

}