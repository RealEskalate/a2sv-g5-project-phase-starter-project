import 'dart:math';

import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_event.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_state.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/pages/HomeChat.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/pages/chat_page.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/detail/detail_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/detail/detail_event.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/detail/detail_state.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_event.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/widgets/bottomnavbar.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/get_user/get_user_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/get_user/get_user_event.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/get_user/get_user_state.dart';
import 'package:ecommerce_app_ca_tdd/locator.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:ecommerce_app_ca_tdd/extra/resusetext.dart';

class DetailsPage extends StatefulWidget {
  final ProductModel item;
  const DetailsPage({required this.item, super.key});

  @override
  _DetailsState createState() => _DetailsState();
}

class _DetailsState extends State<DetailsPage> {
  var list = List<int>.generate(7, (i) => i + 1 + 38);
  int selected = 0;

  @override
  Widget build(BuildContext context) {
    context.read<GetUserBloc>().add(GetUserInfoEvent());
    var on = () {
      context.read<DetailBloc>().add(DeleteProductEvent(widget.item.id));
    };

    return BlocProvider(
      create: (context) => sl.get<ChatBloc>(),
      child: Scaffold(
      bottomNavigationBar: Bottomnavbar(),
        backgroundColor: Theme.of(context).colorScheme.surface,
        floatingActionButtonLocation: FloatingActionButtonLocation.miniStartTop,
        floatingActionButton: Container(
          margin: EdgeInsets.only(left: 20, top: 20),
          height: 40,
          width: 40,
          child: FloatingActionButton(
            shape: CircleBorder(),
            backgroundColor: Colors.white70,
            child: Center(
              child: Icon(
                Icons.arrow_back_ios_new,
                color: Color.fromARGB(255, 63, 81, 243),
                size: 20,
              ),
            ),
            onPressed: () {
              Navigator.pop(context);
            },
          ),
        ),
        body: SafeArea(
          child: BlocConsumer<DetailBloc, DetailState>(
            listener: (context, state) {
              if (state is DetailLoaded) {
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(
                    backgroundColor: Colors.teal[100],
                    content: Row(
                      children: [
                        SizedBox(width: 25),
                        Icon(Icons.thumb_up_rounded, color: Colors.yellow),
                        SizedBox(width: 60),
                        Text(
                          "Successfully Deleted !!!!",
                          style: TextStyle(color: Colors.white),
                        ),
                      ],
                    ),
                  ),
                );
                context.read<HomeBloc>().add(GetProductsEvent());
                Navigator.pushNamed(context, '/home');
              }
            },
            builder: (context, state) {
              return SingleChildScrollView(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Image(
                      height: 286,
                      width: double.infinity,
                      fit: BoxFit.cover,
                      image: NetworkImage(widget.item.imagePath),
                    ),
                    Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 32.0),
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: 
                        [
                          Row(
                              mainAxisAlignment: MainAxisAlignment.spaceBetween,
                              children: [
                                Text(
                                  'product',
                                  style: TextStyle(
                                    color: Colors.grey,
                                    fontFeatures: [FontFeature.enable('palt')],
                                    fontSize: 10,
                                  ),
                                ),
                                Row(
                                  children: [
                                    Icon(Icons.star,
                                        color: Color.fromRGBO(255, 215, 0, 1),
                                        size: 19),
                                    Text('(${4.5})',
                                        style: TextStyle(
                                          color: Colors.grey,
                                          fontSize: 12,
                                        )),
                                  ],
                                ),
                              ],
                            ),
                            SizedBox(height: 4),
                            Row(
                              mainAxisAlignment: MainAxisAlignment.spaceBetween,
                              children: [
                                Text(widget.item.name,
                                    style: TextStyle(
                                      color: Color.fromRGBO(62, 62, 62, 1),
                                      fontWeight: FontWeight.w600,
                                      fontFeatures: [
                                        FontFeature.enable('palt')
                                      ],
                                      fontSize: 24,
                                    )),
                                Text(
                                  '\$${widget.item.price}',
                                  style: TextStyle(
                                      fontSize: 16,
                                      fontWeight: FontWeight.w500),
                                )
                              ],
                            ),

                      //     SizedBox(height: 16),
                      //     Text(
                      //       "Men's Shoes",
                      //       style: GoogleFonts.poppins(
                      //         fontSize: 16,
                      //         fontWeight: FontWeight.w400,
                      //         color: Color.fromARGB(255, 170, 170, 170),
                      //       ),
                      //     ),
                      //     SizedBox(height: 8),
                      //     Text(
                      //       widget.item.name,
                      //       style: GoogleFonts.poppins(
                      //         fontSize: 24,
                      //         fontWeight: FontWeight.w600,
                      //       ),
                      //     ),
                      //     SizedBox(height: 16),
                      //     Row(
                      //       mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      //       children: [
                      //         Row(
                      //           children: [
                      //             Icon(
                      //               Icons.star,
                      //               size: 26,
                      //               color: Color.fromARGB(255, 255, 215, 0),
                      //             ),
                      //             SizedBox(width: 4),
                      //             Text(
                      //               "(4.0)",
                      //               style: GoogleFonts.sora(
                      //                 fontSize: 16,
                      //                 fontWeight: FontWeight.w400,
                      //                 color: Color.fromARGB(255, 170, 170, 170),
                      //               ),
                      //             ),
                      //           ],
                      //         ),
                      //         Text(
                      //           "\$${widget.item.price}",
                      //           style: GoogleFonts.sora(
                      //             fontSize: 16,
                      //             fontWeight: FontWeight.w500,
                      //           ),
                      //         ),
                      //       ],
                      //     ),
                          SizedBox(height: 20),
                          Text(
                            "Sizes:",
                            style: GoogleFonts.poppins(
                              fontWeight: FontWeight.w500,
                              fontSize: 20,
                            ),
                          ),
                          Container(
                            height: 60,
                            child: ListView.builder(
                              itemCount: list.length,
                              scrollDirection: Axis.horizontal,
                              itemBuilder: (context, index) {
                                return GestureDetector(
                                  onTap: () {
                                    setState(() {
                                      selected = index;
                                    });
                                  },
                                  child: Card(
                                    color: selected == index
                                        ? const Color(0xff3F51F3)
                                        : Colors.white,
                                    child: SizedBox(
                                      height: 60,
                                      width: 60,
                                      child: Center(
                                        child: reusableText(
                                          "${list[index]}",
                                          FontWeight.w500,
                                          17,
                                          selected == index
                                              ? Colors.white
                                              : Colors.black,
                                        ),
                                      ),
                                    ),
                                  ),
                                );
                              },
                            ),
                          ),
                          SizedBox(height: 20),
                          Text(
                            widget.item.description,
                            style: GoogleFonts.poppins(
                              color: Color.fromRGBO(102, 102, 102, 1),
                              fontSize: 14,
                              fontWeight: FontWeight.w500,
                            ),
                          ),
                          SizedBox(height: 20),
                          BlocBuilder<GetUserBloc, GetUserState>(
                            builder: (context, state) {
                              if (state is GetUserLoaded && state.user.id==widget.item.sellerId.id){
                              return Row(
                                mainAxisAlignment:
                                    MainAxisAlignment.spaceBetween,
                                children: [
                                  SizedBox(
                                    width: 152,
                                    height: 50,
                                    child: OutlinedButton(
                                      style: OutlinedButton.styleFrom(
                                        shape: RoundedRectangleBorder(
                                          borderRadius: BorderRadius.all(
                                            Radius.circular(10),
                                          ),
                                        ),
                                        side: BorderSide(color: Colors.red),
                                      ),
                                      onPressed: () {
                                        var onPress = () {
                                          context.read<DetailBloc>().add(
                                              DeleteProductEvent(
                                                  widget.item.id));
                                        };
                                        showDialog(
                                          context: context,
                                          builder: (context) => AlertDialog(
                                            title: Text(
                                              "Are you sure you want to delete this product?",
                                              style: GoogleFonts.poppins(
                                                  fontSize: 15),
                                            ),
                                            actions: [
                                              TextButton(
                                                onPressed: () {
                                                  Navigator.pop(context);
                                                },
                                                child: Text("Cancel"),
                                              ),
                                              TextButton(
                                                style: ButtonStyle(
                                                  backgroundColor:
                                                      WidgetStatePropertyAll<
                                                          Color>(Colors.red),
                                                ),
                                                onPressed: onPress,
                                                child: Text("Delete",
                                                style: TextStyle(color: const Color.fromARGB(255, 253, 243, 242)),),
                                              ),
                                            ],
                                          ),
                                        );
                                      },
                                    
                                      child: Text(
                                        "Delete",
                                        style: GoogleFonts.poppins(
                                          fontSize: 14,
                                          fontWeight: FontWeight.w500,
                                          color: Colors.red
                                        ),
                                      ),
                                    ),
                                    
                                  ),
                                  SizedBox(
                                    width: 152,
                                    height: 50,
                                    child: ElevatedButton(
                                      style: ElevatedButton.styleFrom(
                                        shape: RoundedRectangleBorder(
                                          borderRadius: BorderRadius.all(
                                            Radius.circular(10),
                                          ),
                                        ),
                                        side: BorderSide(
                                            color: Color(0xff3F51F3)),
                                        foregroundColor: Colors.white,
                                        backgroundColor: Color(0xff3F51F3),
                                      ),
                                      onPressed: () {
                                        Navigator.pushNamed(
                                          context,
                                          '/update',
                                          arguments: widget.item,
                                        );
                                      },
                                      child: Text(
                                        "Update",
                                        style: GoogleFonts.poppins(
                                          fontSize: 14,
                                          fontWeight: FontWeight.w500,
                                        ),
                                      ),
                                    ),
                                  ),
                                ],
                              );
                              }else if (state is GetUserLoaded && state.user.id!=widget.item.sellerId.id){
                                return Column(
                                  children: [
                                    SizedBox(
                                      height:
                                          MediaQuery.of(context).size.height *
                                              0.15,
                                    ),
                                    Center(
                                        child: SizedBox(
                                      width: 366,
                                      height: 45,
                                      child: ElevatedButton(
                                        style: ElevatedButton.styleFrom(
                                        shape: RoundedRectangleBorder(
                                            borderRadius: BorderRadius.all(
                                                Radius.circular(10))),
                                        side: BorderSide(color: Color(0xff3F51F3)),
                                        // overlayColor: Colors.red,
                                        foregroundColor: Colors.white,
                                        backgroundColor: Color(0xff3F51F3),
                                      ),
                                        onPressed: (){
                                         BlocProvider.of<ChatBloc>(context).add(InitiateChatEvent(widget.item.sellerId.id));
                                         BlocListener<ChatBloc, ChatState>(

                                           listener: (context, state){

                                            print("Chat state: $state");
                                             if (state is ChatInitateLoaded){

                                               Navigator.pushNamed(context, '/chatPage', arguments: state.chat);
                                             }
                                             else if (state is ChatInitateFailure){
                                                showError(context, "Failed to initiate chat");
                                             }else{
                                               print("Chat state: $state");
                                             }
                                           },);
                                          // context.read<ChatBloc>().add(InitiateChatEvent(widget.item.sellerId.id));
                                          
                                        },
                                        child: Row(
                                          children: [
                                            SizedBox(
                                              width: MediaQuery.of(context)
                                                      .size
                                                      .width *
                                                  0.17,
                                            ),
                                            Icon(Icons.phone,
                                                color: Colors.white),
                                            SizedBox(
                                              width: MediaQuery.of(context)
                                                      .size
                                                      .width *
                                                  0.02,
                                            ),
                                            Text(
                                              "Contact Seller",
                                              style: GoogleFonts.poppins(
                                                  fontSize: 14,
                                                  fontWeight: FontWeight.w500),
                                            ),
                                          ],
                                        ),
                                      ),
                                    )),
                                  ],
                                );
                              } else {
                                return SizedBox();
                              }
                            },
                          ),
                        ],
                      ),
                    ),
                  ],
                ),
              );
            },
          ),
        ),
      ),
    );
  }
}
