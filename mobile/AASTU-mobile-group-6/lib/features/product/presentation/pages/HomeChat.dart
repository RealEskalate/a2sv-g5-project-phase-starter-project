// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables
import "package:flutter/material.dart";

class Chat extends StatelessWidget {
  const Chat({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(),
        body: Column(
          children: [
            Text('--------profile page-------'),

            ClipRRect(
              borderRadius: BorderRadius.only(
                topLeft: Radius.circular(50),
                topRight: Radius.circular(50),
              ),
              child: Container(
                  width: MediaQuery.of(context).size.width,
                  height: MediaQuery.of(context).size.height * 0.65,
                  // color: Colors.blue,
                  //
                  child: Center(
                    child: Container(
                        width: MediaQuery.of(context).size.width * 0.95,
                        height: MediaQuery.of(context).size.height * 0.6,
                        color: Colors.white,
                        child: Align(
                          alignment: Alignment.topLeft,
                          child: SingleChildScrollView(
                            child: Column(
                              children: [
                                _duplicate(context,'images/av1.png', 'Estifanos Zinabu', 'How are you today?'),
                                _duplicate(context,'images/av2.png', 'Estifanos Zinabu', 'How are you today?'),
                                _duplicate(context,'images/av3.png', 'Estifanos Zinabu', 'How are you today?'),
                                _duplicate(context,'images/av1.png', 'Estifanos Zinabu', 'How are you today?'),
                                _duplicate(context,'images/av2.png', 'Estifanos Zinabu', 'How are you today?'),
                                _duplicate(context,'images/av3.png', 'Estifanos Zinabu', 'How are you today?'),
                              ],
                            ),
                          ),
                        )),
                  )),
            ),
            Text('----------bottom Bar----------'),
          ],
        ));
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
          radius: 35,
          child: Icon(
              Icons.person,
            size: 35,
          ),
        ),
        Container(
          margin: EdgeInsets.only(left: 5),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            mainAxisAlignment: MainAxisAlignment.end,
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
            mainAxisAlignment: MainAxisAlignment.end,
            children: [
              Text(
                '2 min ago',
                style: TextStyle(
                  color: Colors.grey,
                ),
              ),
              Container(
                // margin: EdgeInsets.only(top: ),
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
