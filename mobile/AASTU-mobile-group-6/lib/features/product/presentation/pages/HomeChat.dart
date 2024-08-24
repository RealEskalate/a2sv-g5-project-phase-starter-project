import 'package:flutter/material.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/widgets/bottomnavbar.dart';

class Chat extends StatelessWidget {
  const Chat({super.key});

  @override
  Widget build(BuildContext context) {
    final List<Map<String, String>> activeUsers = [
      {'name': 'Adil', 'profilePic': 'assets/av1.png'},
      {'name': 'Marina', 'profilePic': 'assets/av2.png'},
      {'name': 'Dean', 'profilePic': 'assets/av3.png'},
      {'name': 'Max', 'profilePic': 'assets/av1.png'},
      {'name': 'Me', 'profilePic': 'assets/av3.png'},
      {'name': 'Adil', 'profilePic': 'assets/av1.png'},
      {'name': 'Marina', 'profilePic': 'assets/av2.png'},
      {'name': 'Dean', 'profilePic': 'assets/av3.png'},
      {'name': 'Max', 'profilePic': 'assets/av1.png'},
      {'name': 'Me', 'profilePic': 'assets/av3.png'},
    ];

    final List<Color> userColors = [
      Colors.red,
      Colors.green,
      Colors.blue,
      Colors.orange,
      Colors.purple,
      Colors.teal,
      Colors.yellow,
      Colors.pink,
      Colors.cyan,
      Colors.amber,
    ];

    return Scaffold(
      appBar: AppBar(
        title: const Text(
          'Search Users',
          style: TextStyle(color: Colors.white),
        ),
        leading: const Icon(
          Icons.search,
          color: Colors.white,
        ),
        backgroundColor: const Color.fromARGB(255, 72, 168, 246),
      ),
      body: Column(
        children: [
          Container(
            margin: EdgeInsets.only(bottom: 10),
            height: 109, // Set a fixed height for the horizontal list
            color: const Color.fromARGB(255, 72, 168, 246),
            child: ListView.builder(
              scrollDirection: Axis.horizontal,
              itemCount: activeUsers.length,
              itemBuilder: (context, index) {
                final user = activeUsers[index];
                return GestureDetector(
                  onTap: () {
                    // Handle user tap
                  },
                  child: Padding(
                    padding: const EdgeInsets.all(8.0),
                    child: Column(
                      children: [
                        CircleAvatar(
                          backgroundImage: AssetImage(user['profilePic']!),
                          radius: 30,
                          backgroundColor:
                              userColors[index % userColors.length],
                        ),
                        const SizedBox(height: 8),
                        Text(
                          user['name']!,
                          style: const TextStyle(color: Colors.white),
                        ),
                      ],
                    ),
                  ),
                );
              },
            ),
          ),
          Expanded(
            child: Container(
              width: MediaQuery.of(context).size.width,
              color: Colors.white,
              child: SingleChildScrollView(
                child: Column(
                  children: [
                    _duplicate(
                      context,
                      'assets/av1.png',
                      'Estifanos Zinabu',
                      'How are you today?',
                    ),
                    _duplicate(
                      context,
                      'assets/av2.png',
                      'Estifanos Zinabu',
                      'How are you today?',
                    ),
                    _duplicate(
                      context,
                      'assets/av3.png',
                      'Estifanos Zinabu',
                      'How are you today?',
                    ),
                    _duplicate(
                      context,
                      'assets/av1.png',
                      'Estifanos Zinabu',
                      'How are you today?',
                    ),
                    _duplicate(
                      context,
                      'assets/av2.png',
                      'Estifanos Zinabu',
                      'How are you today?',
                    ),
                    _duplicate(
                      context,
                      'assets/av3.png',
                      'Estifanos Zinabu',
                      'How are you today?',
                    ),
                  ],
                ),
              ),
            ),
          ),
        ],
      ),
      bottomNavigationBar: Bottomnavbar(),
    );
  }
}

Widget _duplicate(
  BuildContext context,
  String imageUrl,
  String name,
  String peekMessage,
) {
  return Container(
    margin: EdgeInsets.only(bottom: 30),
    width: MediaQuery.of(context).size.width * 0.95,
    height: MediaQuery.of(context).size.height * 0.08,
    child: Row(
      children: [
        CircleAvatar(
          backgroundColor: Colors.yellow,
          radius: 30,
          child: Icon(
            Icons.person,
            size: 35,
          ),
        ),
        Container(
          margin: EdgeInsets.only(left: 10),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Text(
                name,
                style: TextStyle(
                  fontWeight: FontWeight.w800,
                  fontSize: 20,
                ),
              ),
              Text(
                peekMessage,
                style: TextStyle(
                  fontSize: 15,
                  color: Colors.grey,
                ),
              )
            ],
          ),
        ),
        Expanded(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.end,
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Text(
                '2 min ago',
                style: TextStyle(
                  color: Colors.grey,
                ),
              ),
              Container(
                child: CircleAvatar(
                  backgroundColor: Colors.purple,
                  radius: 13,
                  child: Text(
                    '4',
                    style: TextStyle(
                      color: Colors.white60,
                    ),
                  ),
                ),
              ),
            ],
          ),
        ),
      ],
    ),
  );
}
