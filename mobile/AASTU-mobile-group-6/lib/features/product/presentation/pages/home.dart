// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables

import 'package:ecommerce_app_ca_tdd/features/chat/socket_n/chatbox.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/theme_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/HomeChat.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/Theme.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/widgets/bottomnavbar.dart';
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

      return Future.delayed(Duration(seconds: 3));
    }

    const maxNum = 10.0;
    return Scaffold(
      bottomNavigationBar: Container(child: Bottomnavbar()),
      backgroundColor: Theme.of(context).colorScheme.surface,
      appBar: AppBar(
        automaticallyImplyLeading:
            false, //remove pushnamed navigate back button

        title: Container(
          width: MediaQuery.of(context).size.width,
          // color: Colors.white,
          child: Row(
            children: [
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  // -----the box
                  Container(
                    // color: Colors.red,
                    margin: EdgeInsets.only(top: 4),
                    child: SizedBox(
                        width: 70, height: 70, child: ImagePickerIconButton()),
                  ),

                  // -------the day
                  GestureDetector(
                    onTap: () {
                      Navigator.pushNamed(context, '/logout');
                    },
                    child: Container(
                      // color: Colors.purple,
                      padding: EdgeInsets.only(top: 4, left: 10),
                      child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text("July 31, 2024",
                                style: GoogleFonts.syne(
                                    fontWeight: FontWeight.w500, fontSize: 15
                                    // color: Theme.of(context).colorScheme.onSurface,
                                    )),
                            Row(children: [
                              Text("Hello, ",
                                  style: GoogleFonts.sora(
                                      fontWeight: FontWeight.w300,
                                      fontSize: 15,
                                      color:
                                          Color.fromARGB(255, 143, 138, 138))),
                              BlocBuilder<GetUserBloc, GetUserState>(
                                builder: (context, state) {
                                  if (state is GetUserLoading) {
                                    return Text("estif",
                                        style: GoogleFonts.sora(
                                            fontWeight: FontWeight.w600,
                                            fontSize: 15));
                                  } else if (state is GetUserLoaded) {
                                    return Text("${state.user.name}",
                                        style: GoogleFonts.sora(
                                          fontWeight: FontWeight.w600,
                                          fontSize: 15,
                                          // color: Theme.of(context)
                                          //     .colorScheme
                                          //     .onSurface,
                                        ));
                                  } else {
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

              // ------------the last 2 icons--------
              Spacer(),

              Row(
                mainAxisAlignment: MainAxisAlignment.end,
                children: [

                  // --------DARKMODE TOGGLE
                  Container(
                    margin: EdgeInsets.only(right: 15),
                    child: GestureDetector(
                      onTap: () {
                        Provider.of<ThemeProvider>(context, listen: false)
                            .toggleTheme();
                      },
                      child: Container(
                        margin: EdgeInsets.only(top: 7),
                        child: Icon(
                          Provider.of<ThemeProvider>(context).themeData ==
                                  darkmode
                              ? Icons.wb_sunny // Sun icon for light mode
                              : Icons.nights_stay, // Moon icon for dark mode
                          size: 30,
                          // color: Theme.of(context).colorScheme.onSurface,
                        ),
                      ),
                    ),
                  ),

                  // -----------chat icon------
                  Container(
                      // color: Colors.yellow,
                      child: GestureDetector(
                          onTap: () {
                            SocketService().connectToServer();
                            Navigator.pushNamed(context, '/HomeChat');
                          },
                          child: Transform.rotate(
                            angle: 5.5,
                            child: Icon(
                              Icons.send_rounded,
                              size: 30,
                            ),
                          ))),
                ],
              ),
            ],
          ),
        ),
      ),
      body: SafeArea(
          child: Padding(
        padding: const EdgeInsets.only(left: 21.0, right: 21.0),
        child: Column(
          children: [
            Container(
              height: MediaQuery.of(context).size.height * 0.052,
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Text(
                    "Available Products",
                    style: TextStyle(
                        fontFamily: "Poppins",
                        color: Theme.of(context).colorScheme.onSurface,
                        fontSize: 24,
                        fontWeight: FontWeight.w600),
                  ),
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
                            color: Theme.of(context).colorScheme.onSurface,
                            onRefresh: _refresh,
                            child: ListView.builder(
                              itemCount: state.products.length,
                              itemBuilder: (context, index) {
                                return GestureDetector(
                                    onTap: () {
                                      Navigator.pushNamed(
                                        context,
                                        '/detail',
                                        arguments: state.products[index],
                                      );
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
        // color: Colors.red,
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
                          color: Theme.of(context).colorScheme.onSurface,
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
                            color: Theme.of(context).colorScheme.onSurface,
                          )),
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
                            color: Theme.of(context).colorScheme.onSurface,
                            fontSize: 14,
                            fontWeight: FontWeight.w500),
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
                                  color:
                                      Theme.of(context).colorScheme.onSurface)),
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
  if (result != null) {
    return false;
  } else {
    return true;
  }
}
