// import 'dart:ffi';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../service_locator.dart';
import '../bloc/getallproductbloc/bloc/product_bloc.dart';
import 'homepage.dart';
import 'package:flutter/material.dart';

class Searchpage extends StatefulWidget {
  Searchpage({super.key});
  bool filter = true;
  @override
  State<Searchpage> createState() => _Searchpage();
}

class _Searchpage extends State<Searchpage> {
  RangeValues current = const RangeValues(20, 80);
  TextEditingController searchController = TextEditingController();
  var product_bloc = getIt<ProductBloc>();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        backgroundColor: Colors.white,
        body: Padding(
          padding: EdgeInsets.only(left: 20.0, top: 50, right: 20.0),
          child: Column(
            children: [
              Row(
                children: [
                  IconButton(
                    icon: Icon(Icons.arrow_back_ios, color: Color(0xFF3f51f3)),
                    onPressed: () {
                      Navigator.pop(context);
                    },
                  ),
                  const SizedBox(width: 60),
                  const Text(
                    'Search Products',
                    style: TextStyle(
                      fontSize: 20,
                      // fontWeight: FontWeight.bold,
                    ),
                  ),
                ],
              ),
              SizedBox(
                height: 20,
              ),
              Row(
                children: [
                  Expanded(
                    child: Stack(
                      children: [
                        TextField(
                          controller: searchController,
                          decoration: InputDecoration(
                              labelText: 'Search',
                              border: OutlineInputBorder(
                                  borderRadius:
                                      BorderRadius.all(Radius.circular(10)))),
                        ),
                        Positioned(
                          top: 5,
                          right: 0,
                          child: IconButton(
                              onPressed: () {
                                product_bloc.add(
                                    GetProductEvent(searchController.text));
                              },
                              icon: Icon(
                                Icons.arrow_forward,
                                color: Color(0xFF3f51f3),
                              )),
                        )
                      ],
                    ),
                  ),
                  SizedBox(width: 6),
                  SizedBox(
                    height: 60,
                    width: 50,
                    child: Material(
                      color: Color(0xFF3f51f3),
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.all(Radius.circular(10)),
                      ),
                      child: IconButton(
                        onPressed: () {
                          setState(() {
                            widget.filter = !widget.filter;
                          });
                        },
                        icon: Icon(
                          Icons.filter_list,
                          color: Colors.white,
                        ),
                      ),
                    ),
                  )
                ],
              ),
              SizedBox(
                height: 15,
              ),
              Expanded(
                child: BlocBuilder<ProductBloc, ProductState>(
                  builder: (context, state) {
                    if (state is getloaded){
                    return SingleChildScrollView(
                      child: Column(
                        children: [
                          Container(
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
                                  Image(
                                    image: NetworkImage(state.product.image),
                                  ),
                                  SizedBox(height: 8.0),
                                  Row(
                                    mainAxisAlignment:
                                        MainAxisAlignment.spaceBetween,
                                    children: [
                                      Padding(
                                        padding: EdgeInsets.only(left: 14.0),
                                        child: Text(
                                          state.product.name,
                                          style: TextStyle(
                                            fontSize: 20,
                                            fontWeight: FontWeight.bold,
                                          ),
                                        ),
                                      ),
                                      Padding(
                                        padding: EdgeInsets.only(right: 14.0),
                                        child: Text(
                                          state.product.price.toString(),
                                          style: TextStyle(
                                            fontWeight: FontWeight.bold,
                                          ),
                                        ),
                                      ),
                                    ],
                                  ),
                                  SizedBox(height: 4.0),
                                  Row(
                                    mainAxisAlignment:
                                        MainAxisAlignment.spaceBetween,
                                    children: [
                                      Padding(
                                        padding: EdgeInsets.only(left: 14.0),
                                        child: Text(
                                          state.product.description,
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
                                            Text('4.0'),
                                          ],
                                        ),
                                      ),
                                    ],
                                  ),
                                ],
                              ),
                            ),
                          ),
                        ],
                      ),
                    );
                  }
                  else if(state is getloading){
                    return Center(child: CircularProgressIndicator(),);
                  }
                  else{
                    return Center(child: Text("Please search any product"),);
                  }
                  },
                ),
              ),
              widget.filter
                  ? Container(
                      width: double.infinity,
                      height: 300,
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          SizedBox(
                            height: 10,
                          ),
                          Text(
                            "Price",
                            style: TextStyle(
                              fontSize: 18,
                            ),
                          ),
                          RangeSlider(
                              min: 0,
                              max: 100,
                              activeColor: Color(0xFF3f51f3),
                              inactiveColor: Colors.grey,
                              values: current,
                              onChanged: (RangeValues values) {
                                setState(() {
                                  current = values;
                                });
                              }),
                          SizedBox(
                            height: 35,
                          ),
                          Container(
                            width: double.infinity,
                            height: 45,
                            child: Center(
                              child: Text(
                                'ADD',
                                style: TextStyle(color: Colors.white),
                              ),
                            ),
                            decoration: BoxDecoration(
                              color: Color(0xFF3f51f3),
                              borderRadius:
                                  BorderRadius.all(Radius.circular(10)),
                            ),
                          ),
                        ],
                      ),
                    )
                  : Text(''),
            ],
          ),
        ));
  }
}
