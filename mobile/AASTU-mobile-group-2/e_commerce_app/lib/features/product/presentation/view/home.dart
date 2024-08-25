import 'package:e_commerce_app/features/auth/presentation/bloc/auth_bloc.dart';
import 'package:e_commerce_app/features/auth/presentation/bloc/auth_state.dart';
import 'package:e_commerce_app/features/product/domain/entities/product.dart';
import 'package:e_commerce_app/features/product/presentation/bloc/home/home_bloc.dart';
import 'package:e_commerce_app/features/product/presentation/bloc/search/search_product_bloc.dart';
import 'package:e_commerce_app/features/product/presentation/view/widgets/widgets.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:provider/provider.dart';

import '../../../auth/presentation/bloc/auth_event.dart';

class Home extends StatelessWidget {
  Home({super.key});

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        backgroundColor: Colors.white,
        body: Container(
          color: Colors.white,
          padding: EdgeInsets.all(20),
          child: Column(
            children: [
              Row(
                children: [
                  Container(
                    height: 50,
                    width: 50,
                    decoration: BoxDecoration(
                        color: Color.fromARGB(255, 204, 204, 204),
                        borderRadius: BorderRadius.circular(11)),
                  ),
                  SizedBox(
                    width: 15,
                  ),
                  Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        "July 14,2023",
                        style: TextStyle(
                            color: Color.fromARGB(255, 170, 170, 170),
                            fontWeight: FontWeight.w500),
                      ),
                      Row(
                        children: [
                          Text(
                            "Hello ,",
                            style: TextStyle(
                                fontSize: 15, fontWeight: FontWeight.w400),
                          ),
                          BlocConsumer<AuthBloc, AuthState>(
                            listener: (context, state) {
                              // print(state);

                              if (state is LoginSuccess) {
                                context.read<AuthBloc>().add(GetUserEvent());
                              }
                            },
                            builder: (BuildContext context, AuthState state) {
                              if (state is GetUserSuccess) {
                                return Text(
                                  state.name,
                                  style: TextStyle(
                                    fontWeight: FontWeight.bold,
                                    fontSize: 15,
                                  ),
                                );
                              } else {
                                return Text(
                                  "",
                                  style: TextStyle(
                                    fontWeight: FontWeight.bold,
                                    fontSize: 15,
                                  ),
                                );
                              }
                            },
                          )
                        ],
                      )
                    ],
                  ),
                  Spacer(),
                  OutlinedButton(
                    onPressed: () {},
                    child: Icon(
                      Icons.notifications_none,
                      color: Colors.black,
                    ),
                    style: OutlinedButton.styleFrom(
                        padding: EdgeInsets.all(1),
                        minimumSize: Size(40, 40),
                        shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(9)),
                        side: BorderSide(
                            color: Color.fromARGB(255, 217, 217, 217)),
                        backgroundColor: Colors.white),
                  )
                  // ButtonIcon(icon: Icons.notifications_none_outlined,background: Colors.white,color: Colors.black,)
                ],
              ),
              const SizedBox(
                height: 15,
              ),
              // product section

              Flexible(
                child: Column(
                  children: [
                    Row(
                      children: [
                        Text(
                          "Available Products",
                          style: TextStyle(
                            fontSize: 22,
                            fontWeight: FontWeight.w600,
                          ),
                        ),
                        Spacer(),
                        BlocBuilder<HomeBloc, HomeState>(
                          builder: (context, state) {
                            List<ProductEntity> products = [];
                            if (state is HomeSuccessLoading) {
                              products = state.allProducts;
                            }
                            return OutlinedButton(
                              onPressed: () {
                                context
                                    .read<SearchBloc>()
                                    .add(SearOpened(allProducts: products));
                                Navigator.pushNamed(context, '/searchpage');
                              },
                              child: Icon(
                                Icons.search,
                                color: Color.fromARGB(255, 217, 217, 217),
                              ),
                              style: OutlinedButton.styleFrom(
                                  padding: EdgeInsets.all(2),
                                  minimumSize: Size(40, 40),
                                  shape: RoundedRectangleBorder(
                                      borderRadius: BorderRadius.circular(9)),
                                  side: BorderSide(
                                      color:
                                          Color.fromARGB(255, 217, 217, 217)),
                                  backgroundColor: Colors.white),
                            );
                          },
                        ),
                      ],
                    ),
                    const SizedBox(
                      height: 15,
                    ),

                    BlocBuilder<HomeBloc, HomeState>(builder: (context, state) {
                      if (state is HomeSuccessLoading) {
                        List<ProductEntity> products = state.allProducts;

                        List<Widget> allCards = products.map((Product) {
                          return ItemCard(product: Product);
                        }).toList();
                        return Expanded(
                          child: ListView(
                            children: allCards,
                          ),
                        );
                      } else if (state is HomeProductLoading) {
                        return Center(child: CircularProgressIndicator());
                      } else {
                        return Text("error ");
                      }
                    })
                    //  ListView(
                    //   children: allCards,
                    // )
                    // )
                  ],
                ),
              )
            ],
          ),
        ),
        floatingActionButton: Builder(builder: (context) {
          return FloatingActionButton(
            backgroundColor: Colors.blue,
            onPressed: () {
              Navigator.pushNamed(context, '/insertitem');
            },
            child: Icon(
              Icons.add,
              color: Colors.white,
              size: 36,
            ),
            shape: CircleBorder(),
          );
        }
        ),
      ),
    );
  }
}
