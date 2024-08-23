// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables

import 'package:ecommerce_app_ca_tdd/extra/dark_mode.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_access.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/get_user/get_user_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/get_user/get_user_event.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/get_user/get_user_state.dart';
import 'package:flutter/material.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/add_update.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/details.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:image_picker/image_picker.dart';
import 'package:ecommerce_app_ca_tdd/extra/icon_img.dart';
import 'package:page_transition/page_transition.dart';
import 'package:ecommerce_app_ca_tdd/models/product.dart';
import 'package:ecommerce_app_ca_tdd/extra/overflow_card.dart';
import 'package:provider/provider.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/search.dart';
import 'package:ecommerce_app_ca_tdd/extra/resusetext.dart';
import 'package:ecommerce_app_ca_tdd/models/ext_product.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../../domain/entities/product_entity.dart';
import '../bloc/home_bloc.dart';
import '../bloc/home_event.dart';
import '../bloc/home_state.dart';

class HomePage extends StatefulWidget {
  // final UserAccess token;
  // HomePage({required this.token});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  
  


  @override
  Widget build(BuildContext context) {
    Future<void> _refresh() {
      context.read<HomeBloc>().add(GetProductsEvent());
      context.read<GetUserBloc>().add(GetUserInfoEvent());
    
  
  return  Future.delayed(Duration(seconds: 3));}
    
    const maxNum = 10.0;
    return Scaffold(
      backgroundColor: Colors.white,
      floatingActionButton: SizedBox(
        width: 72,
        height: 72,
        child: FloatingActionButton(
          shape: CircleBorder(),
          backgroundColor: Color.fromARGB(255, 63, 81, 243),
          child: Icon(Icons.add_rounded, color: Colors.white, size: 55),
          // Named ROute Implementation
          onPressed: () {
            Navigator.pushNamed(context, '/add');
          },
        ),
      ),
      body: SafeArea(
          // ignore: prefer_const_constructors
          child: Padding(
        padding: const EdgeInsets.only(left: 21.0, right: 21.0),
        child: Column(
          children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Container(
                      margin: EdgeInsets.only(top: 4),
                      child: SizedBox(
                          width: 80,
                          height: 80,
                          child: ImagePickerIconButton()),
                    ),
                    GestureDetector(
                      onTap: () {
                        Navigator.pushNamed(context,'/logout');
                      },
                      child: Container(
                        padding: EdgeInsets.only(top: 4, left: 10),
                        child: Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              Text(
                                "July 31, 2024",
                                style: GoogleFonts.syne(
                                    fontWeight: FontWeight.w500,
                                    color: Color.fromARGB(255, 170, 170, 170)),
                              ),
                              Row(children: [
                                Text("Hello,",
                                    style: GoogleFonts.sora(
                                        fontWeight: FontWeight.w400,
                                        color:
                                            Color.fromARGB(255, 102, 102, 102))),
                                BlocBuilder<GetUserBloc, GetUserState>(
                                  builder: (context, state) {
                                    if (state is GetUserLoading){
                                      return Text("Fetching User...",
                                        style: GoogleFonts.sora(
                                            fontWeight: FontWeight.w600));
                                    }else if (state is GetUserLoaded){
                                      return Text("${state.user.name}",
                                        style: GoogleFonts.sora(
                                            fontWeight: FontWeight.w600));
                                    }else{
                                      return Text('name');
                                    }
                                    
                                  },
                                ),
                              ])
                            ]),
                      ),
                    ),
                  ],
                ),
                Row(
                  children: [
                    const DarkMode(),
                    Container(
                        decoration: BoxDecoration(
                            border: Border.all(
                                color: Color.fromRGBO(221, 221, 221, 1), width: 2),
                            borderRadius: BorderRadius.circular(9)),
                              child: GestureDetector(
                                  onTap: (){
                                      showDialog(
                                          context: context, 
                                          builder: (context)=> AlertDialog(
                                            title: Text("Are you sure you want to logout ?",style: GoogleFonts.poppins(fontSize: 15),),
                                            actions: [
                                              TextButton(onPressed: (){
                                                Navigator.pop(context);
                                              }, child: Text("Cancel")),
                                              TextButton(
                                                onPressed: (){
                                                  logOut();
                                                  Navigator.pushNamedAndRemoveUntil(context, '/login', (route) => true);
                                                },
                                                child: Text("Log-Out")
                                                )
                                            ],
                                                                      )
                                                                    );
                                  },
                          child: Image(
                              width: 40,
                              height: 40,
                              image: AssetImage(
                                  'assets/icons8-notification-bell-24.png')),
                        )),
                  ],
                )
              ],
            ),

            Container(
              height: MediaQuery.of(context).size.height * 0.052,
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Text(
                    "Available Products",
                    style: TextStyle(
                        fontFamily: "Poppins",
                        fontSize: 24,
                        fontWeight: FontWeight.w600),
                  ),
                  Container(
                      decoration: BoxDecoration(
                          border: Border.all(color: Colors.grey, width: 1),
                          borderRadius: BorderRadius.circular(10)),
                      padding: EdgeInsets.only(right: 2),
                      child: IconButton(
                          onPressed: () {
                            Navigator.push(
                              context,
                              PageTransition(
                                  alignment: Alignment.bottomCenter,
                                  curve: Curves.easeInOut,
                                  duration: Duration(milliseconds: 1200),
                                  reverseDuration: Duration(milliseconds: 400),
                                  type: PageTransitionType.leftToRightPop,
                                  child: searchPage(),
                                  childCurrent: HomePage()),
                            );
                          },
                          icon: Icon(Icons.search),
                          color: Color.fromARGB(255, 217, 217, 217),
                          iconSize: 29))
                ],
              ),
            ),
            // Products
            SizedBox(
              height: 11,
            ),

            BlocBuilder<HomeBloc, HomeState>(
              builder: (context, state) {
                if (state is HomeLoading) {
                  return Center(child: CircularProgressIndicator());
                } else if (state is HomeFailure) {
                  return SnackBar(
                    content: Text(state.message),
                  );
                } else if (state is HomeLoaded) {
                  return Expanded(
                    child: SizedBox(
                      child: SingleChildScrollView(
                        child: SizedBox(
                          height: MediaQuery.of(context).size.height * 0.8,
                          child: RefreshIndicator(
                            onRefresh:_refresh,
                            child: ListView.builder(
                              itemCount: state.products.length,
                              itemBuilder: (context, index) {
                                return GestureDetector(
                                    onTap: () {
                                      Navigator.pushNamed(context, '/detail',
                                          arguments: state.products[index]);
                                    },
                                    child: OverflowCard(
                                      product: state.products[index],
                                    ));
                              },
                            ),
                          ),
                        ),
                      ),
                    ),
                  );
                }
                return Container(); // Add a return statement at the end
              },
            ),

            //  Button Blue

            // onPressed: onPressed)
          ],
        ),
      )),
    );
  }
}

Future<String?> getAccessToken() async {
  SharedPreferences prefs = await SharedPreferences.getInstance();
  var token = prefs.getString('access_token');
  return token.toString();
}




class OverflowCard extends StatelessWidget {
  final ProductModel product;

  const OverflowCard({required this.product});

  @override
  Widget build(BuildContext context) {
    return InkWell(
      child: Card(
        elevation: 2.0,
        margin: const EdgeInsets.all(8.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            ClipRRect(
              borderRadius: BorderRadius.only(
                  topLeft: Radius.circular(16.0),
                  topRight: Radius.circular(16.0)),
              child: Image(
                fit: BoxFit.cover,
                width: MediaQuery.of(context).size.width * 1,
                height: MediaQuery.of(context).size.height * 0.3,
                image: NetworkImage(product.imagePath), // For local image
              ),
            ), // For network image
            SizedBox(height: 5.0),
            Padding(
              padding: const EdgeInsets.only(
                  left: 16.0, top: 4, bottom: 12, right: 21),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        product.name,
                        style: GoogleFonts.poppins(
                          fontSize: 20.0,
                          fontWeight: FontWeight.w500,
                        ),
                      ),
                      SizedBox(height: 8.0),
                      Text('Ecommerce Items',
                      overflow: TextOverflow.ellipsis,
                          style: GoogleFonts.poppins(
                              fontSize: 15.0,
                              fontWeight: FontWeight.w400,
                              color: Color.fromRGBO(217, 217, 217, 1))),
                    ],
                  ),
                  SizedBox(height: 8.0),
                  Column(
                    mainAxisAlignment: MainAxisAlignment.start,
                    crossAxisAlignment: CrossAxisAlignment.end,
                    children: [
                      Text(
                        '\$' + product.price.toString(),
                        
                        overflow: TextOverflow.fade,
                        style: GoogleFonts.sora(
                            fontSize: 14, fontWeight: FontWeight.w500),
                      ),
                      SizedBox(height: 8.0),
                      Row(
                        children: [
                          Icon(
                            Icons.star,
                            size: 24,
                            color: Color.fromARGB(255, 255, 215, 0),
                          ),
                          Text("(4.0)",
                              style: GoogleFonts.sora(
                                  fontSize: 12,
                                  fontWeight: FontWeight.w400,
                                  color: Color.fromARGB(255, 170, 170, 170))),
                        ],
                      )
                    ],
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}

// end of home page

class CustomBlocObserver extends BlocObserver {
  @override
  void onChange(BlocBase bloc, Change change) {
    super.onChange(bloc, change);
    print('Current State: ${change.currentState}');
    print('Next State: ${change.nextState}');
  }
}

Future<void> logOut() async {
  SharedPreferences remove = await SharedPreferences.getInstance();
  remove.remove('access_token');
}
Future<bool?> checkToken() async {
  SharedPreferences remove = await SharedPreferences.getInstance();
  var result = remove.getString('access_token');
  if (result != null){
    return false;
  }else{
  return true;

  }
}