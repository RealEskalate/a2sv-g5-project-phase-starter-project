
import 'package:flutter/material.dart';

class ChatListPage extends StatefulWidget {
  @override
  _ChatListPageState createState() => _ChatListPageState();
}

class _ChatListPageState extends State<ChatListPage> {
 List<String> names = ['Aschalew Abayneh','Sari Amin','Samuel Tollosa','sdds','sdsdsdsa','sdsaas','dsd','sssds','sdds','sdsdsdsa','sdsaas','dsd','sssds','sdds','sdsdsdsa','sdsaas','dsd','sssds','sdds','sdsdsdsa'];
List<String> Messages = ['hello, how are you? ','aa','sssds','sdds','sdsdsdsa','sdsaas','dsd','sssds','sdds','sdsdsdsa','aa','aa','sssds','sdds','sdsdsdsa','sdsaas','dsd','sssds','sdds','sdsdsdsa'];

  final List<String> profilePhotos = [
    'assets/image/pro1.jpg', 'assets/image/pro2.jpg','assets/image/pro3.jpg',   'assets/image/pro4.jpg',   
    'assets/image/grouppro.png','assets/image/grouppro2.png', 'assets/image/pro5.png',  'assets/image/pro6.jpg','assets/image/pro7.jpg',
    ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.blue,
      appBar: AppBar(
        backgroundColor: Colors.blue,
        leading: IconButton(
          icon: const Icon(Icons.arrow_back_ios),
          onPressed: () {
            // Handle the search icon press
            Navigator.pop(context);
          },
        ),
      ),
      body: Column(
        children: [
         Container(
           padding: const EdgeInsets.symmetric(horizontal: 3),
           width: MediaQuery.of(context).size.width,
           height: 150,
           color: Colors.blue,
           child: ListView(
             scrollDirection: Axis.horizontal,
             children: [
               Mystory('My status', 'assets/image/pro.png'),
               story('Marina', 'assets/image/pro1.jpg'),
               story('Dean', 'assets/image/pro2.jpg'),
               story('Max', 'assets/image/pro3.jpg'),
               story('My status', 'assets/image/pro4.jpg'),
               story('Adil', 'assets/image/pro5.png'),
               story('Marina', 'assets/image/pro6.jpg'),
               story('Dean', 'assets/image/pro7.jpg'),
               story('My status', 'assets/image/pro.jpg'),
               story('Adil', 'assets/image/grouppro.png'),
               story('Marina', 'assets/image/grouppro2.png'),
             ],
           ),
         ),



          
          Expanded(
            child: Container(
             
              decoration: BoxDecoration(
              color: Colors.white,
              borderRadius: BorderRadius.circular(30)
            ),
              child: ListView.builder(
                itemCount: names.length,
                itemBuilder: (context, index) {
                  return GestureDetector(
                    onTap: (){
                      Navigator.pushNamed(context, '/chat-message');
                    },
                    child: chatTileWithReadMessages(
                      names[index],
                      Messages[index],
                      '02:20',
                      index,
                    ),
                  );
                },
              ),
            ),
          ),
        ],
      ),
    );
  }

  Widget chatTileWithReadMessages(String name, String message, String time, int index) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 5),
      child: ListTile(
        leading: CircleAvatar(
          backgroundImage: AssetImage(profilePhotos[index % profilePhotos.length]),
          radius: 30,
        ),
        title: Text(
          name,
          style: const TextStyle(
            fontWeight: FontWeight.bold,
          ),
        ),
        subtitle: Text(message),
        trailing: Text(time, style: TextStyle(fontSize: 12, color: Colors.grey)),
      ),
    );
  }

  Widget Mystory(String name, String image) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Column(
        children: [
          Stack(
            children: [
              Container(
                width: 60,
                height: 60,
                decoration: BoxDecoration(
                  shape: BoxShape.circle,
                  border: Border.all(
                    color: Colors.blue,
                    width: 3,
                  ),
                ),
                child: CircleAvatar(
                  backgroundImage: AssetImage(image),
                ),
              ),
              Positioned(
                bottom: 0,
                right: 0,
                child: Container(
                  width: 20,
                  height: 20,
                  decoration: const BoxDecoration(
                    color: Colors.blue,
                    shape: BoxShape.circle,
                  ),
                  child: const Icon(
                    Icons.add,
                    size: 14,
                    color: Colors.white,
                  ),
                ),
              ),
            ],
          ),
          const SizedBox(height: 8),
          Text(
            name,
            style: const TextStyle(
              color: Colors.white,
              fontWeight: FontWeight.bold,
              fontSize: 12,
            ),
          ),
        ],
      ),
    );
  }

  Widget story(String name, String image) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Column(
        children: [
          Container(
            width: 60,
            height: 60,
            decoration: BoxDecoration(
              shape: BoxShape.circle,
              border: Border.all(
                color: Colors.blue,
                width: 3,
              ),
            ),
            child: CircleAvatar(
              backgroundImage: AssetImage(image),
            ),
          ),
          const SizedBox(height: 8),
          Text(
            name,
            style: const TextStyle(
              color: Colors.white,
              fontSize: 12,
            ),
          ),
        ],
      ),
    );
  }
}