import 'package:ecommerce/features/auth/data/model/user_model.dart';
import 'package:ecommerce/features/auth/presentation/login.dart';
import 'package:ecommerce/features/auth/presentation/register.dart';
import 'package:flutter/material.dart';
import 'package:get_it/get_it.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../../../../service_locator.dart';

import 'package:dartz/dartz.dart' as dartz;

import 'package:ecommerce/features/product/presentation/pages/details_page.dart';
import 'package:ecommerce/features/product/presentation/pages/search.dart';
import 'package:flutter/cupertino.dart';

import 'package:flutter/widgets.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../../../service_locator.dart';
import '../../domain/entities/product.dart';
import '../../domain/usecases/addproduct.dart';
import '../../domain/usecases/getallproduct.dart';

import '../bloc/getallproductbloc/bloc/product_bloc.dart';
import 'addupdate.dart';

class MyHomePage extends StatefulWidget {
  final UserModel user;
  MyHomePage({super.key, required this.title, required this.user});

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  late GetAllProductUsecase getAllUseCase;
  late AddProductUsecase addProductUsecase;
  var product;
  void initState() {
    super.initState();
    print("homepage user");
    print(widget.user);
    _initializeData(); // Initialize data asynchronously
  }

  Future<void> _initializeData() async {
    getAllUseCase = getIt<GetAllProductUsecase>();
    addProductUsecase = getIt<AddProductUsecase>();
    product = await getAllUseCase.getall();
    setState(() {
      product = product;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        automaticallyImplyLeading: false,
        backgroundColor: Colors.white,
        title: Row(
          children: [
            ElevatedButton(
                style: ElevatedButton.styleFrom(
                    minimumSize: const Size(55, 55),
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(15),
                    )),
                onPressed: null,
                child: const Text('')),
            const SizedBox(width: 20),
            Expanded(
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      const Text(
                        'July 14,2023',
                        style: TextStyle(
                            fontSize: 12,
                            color: Colors.grey,
                            fontFamily: 'RobotoMono'),
                      ),
                      Row(
                        children: [
                          const Text(
                            'Hello,',
                            style: TextStyle(
                                fontSize: 15,
                                color: Colors.grey,
                                fontFamily: 'RobotoMono'),
                          ),
                          Text(
                            widget.user.name,
                            style: GoogleFonts.lato(
                              fontSize: 15,
                            ),
                          )
                        ],
                      ),
                    ],
                  ),
                  OutlinedButton(
                      onPressed: () async {
                        var sharedPreference =
                            await SharedPreferences.getInstance();
                        sharedPreference.setString('token', '');
                        Navigator.push(
                          context,
                          MaterialPageRoute(builder: (context) => login()),
                        );
                      },
                      iconAlignment: IconAlignment.start,
                      style: OutlinedButton.styleFrom(
                          backgroundColor:
                              const Color.fromARGB(255, 255, 255, 255),
                          fixedSize: const Size(95, 40),
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(15),
                          )),
                      child: Text(
                        'Logout',
                      ))
                ],
              ),
            ),
          ],
        ),
      ),
      body: BlocProvider(
        create: (context) => getIt<ProductBloc>()..add(GetAllProductEvent()),
        child: SingleChildScrollView(
          child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20.0),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children: <Widget>[
                const SizedBox(
                  height: 20,
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    const Text(
                      'Available Products',
                      style: TextStyle(
                        fontSize: 20,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                    OutlinedButton(
                        style: OutlinedButton.styleFrom(
                          backgroundColor: Colors.white,
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(15),
                          ),
                          fixedSize: Size(50, 25),
                        ),
                        onPressed: () {
                          // getIt<ProductBloc>().close();
                          Navigator.push(
                            context,
                            MaterialPageRoute(
                                builder: (context) => Searchpage()),
                          );
                        },
                        child: const Icon(
                          Icons.search,
                          color: Colors.black,
                        ))
                  ],
                ),
                SizedBox(
                  height: 15,
                ),
                RefreshIndicator(
                  onRefresh: () => refresh(context),
                  child: Container(
                    height: MediaQuery.of(context).size.height,
                    child: BlocBuilder<ProductBloc, ProductState>(
                        builder: (context, state) {
                      // print('Rediet');
                      // print(state);
                      if (state is homeloading) {
                        return const Center(
                          child: CircularProgressIndicator(),
                        );
                      } else if (state is homeloaded) {
                        return ListView.builder(
                          itemCount: state.product.length,
                          itemBuilder: (context, index) {
                            return Column(
                              children: [
                                ProductCard(
                                    user: widget.user,
                                    sampleproduct: Productentity(
                                        id: state.product[index].id,
                                        image: state.product[index].image,
                                        name: state.product[index].name,
                                        price: state.product[index].price,
                                        description:
                                            state.product[index].description,
                                        seller: state.product[index].seller)),
                                const SizedBox(
                                  height: 20,
                                ),
                              ],
                            );
                          },
                        );
                      } else {
                        return Center(
                          child: Text("Failed fetching"),
                        );
                      }
                    }),
                  ),
                )
              ],
            ),
          ),
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          Navigator.push(
            context,
            MaterialPageRoute(
                builder: (context) => AddUpdate(user: widget.user)),
          );
        },
        tooltip: 'Increment',
        child: Icon(
          Icons.add,
          color: Colors.white,
        ),
        backgroundColor: Color(0xFF3f51f3),
        // Using the hex color
        shape: RoundedRectangleBorder(
          borderRadius:
              BorderRadius.circular(50), // Adjust the radius as needed
        ),
      ), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}

class ProductCard extends StatelessWidget {
  const ProductCard(
      {super.key, required this.sampleproduct, required this.user});
  final Productentity sampleproduct;
  final UserModel user;
  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        Navigator.push(
            context,
            MaterialPageRoute(
                builder: (context) => DetailsPage(
                      user: user,
                      sampleproduct: sampleproduct,
                    )));
      },
      child: Container(
        decoration: BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.circular(12.0),
          boxShadow: const [
            BoxShadow(
              color: Colors.black12,
              blurRadius: 8.0,
              offset: Offset(0, 4),
            ),
          ],
        ),
        child: Padding(
          padding: EdgeInsets.all(0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              // Text("imageUrl : ${sampleproduct.image}"),
              Image(
                image: NetworkImage(sampleproduct.image),
              ),
              const SizedBox(height: 8.0),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Padding(
                    padding: const EdgeInsets.only(left: 14.0),
                    child: Text(
                      sampleproduct.name,
                      style: const TextStyle(
                        fontSize: 20,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                  ),
                  Padding(
                    padding: EdgeInsets.only(right: 14.0),
                    child: Text(
                      sampleproduct.price.toString(),
                      style: const TextStyle(
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                  ),
                ],
              ),
              const SizedBox(height: 4.0),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Padding(
                    padding: EdgeInsets.only(left: 14.0),
                    child: Text(
                      sampleproduct.description,
                      style: TextStyle(
                        color: Colors.grey,
                      ),
                    ),
                  ),
                  Padding(
                    padding: EdgeInsets.only(right: 14.0),
                    child: Row(
                      children: [
                        Icon(
                          Icons.star,
                          color: Colors.amber,
                        ),
                        SizedBox(width: 4.0),
                        // Text(sampleproduct.rating),
                      ],
                    ),
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}

// class Product {
//   final String image;
//   String productname;
//   double price;
//   String description;
//   Product(
//       {required this.image,
//       required this.productname,
//       required this.price,
//       required this.description});
// }

Future<void> refresh(BuildContext context) async {
  BlocProvider.of<ProductBloc>(context)
      .add(GetAllProductEvent()); // await Future.delayed(Duration(seconds: 1));
}
